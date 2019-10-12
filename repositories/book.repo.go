package repositories

import (
	"fmt"
	models "libraryapi/models"

	"github.com/jinzhu/gorm"
)

func FindAllBook(db *gorm.DB) ([]models.Book, error) {
	books := []models.Book{}

	// if err := db.Find(&books).Error; err != nil {
	if err := db.Preload("Category").Preload("Bookshelf").Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func FindBookByID(db *gorm.DB, id uint) ([]models.Book, error) {
	books := []models.Book{}

	// if err := db.Find(&books).Error; err != nil {
	if err := db.Preload("Category").Preload("Bookshelf").Find(&books, id).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func InsertBook(db *gorm.DB, bk *models.Book) (uint, error) {
	if err := db.Create(&bk).Error; err != nil {
		fmt.Println("ada error di insert book :V \n%v", err)
		return 0, err
	}

	return bk.ID, nil
}

func UpdateBook(db *gorm.DB, book *models.Book) (uint, error) {
	bookDB := models.Book{}

	fmt.Println("Obj Users :", book)

	// if err := db.Where("id = ?", book.ID).First(book).Error; err != nil {
	// 	fmt.Println("User tidak ditemukan\ntidak ada book yang a m a n")
	// 	return 0, err
	// }

	if db.Where("id = ?", book.ID).First(&bookDB).RecordNotFound() {
		fmt.Println("Book tidak ditemukan\ntidak ada book yang a m a n")
		return 0, nil
	}

	//set data book dari parameter yang dikirim
	bookDB.ISBN = book.ISBN
	bookDB.Title = book.Title
	bookDB.Author = book.Author
	bookDB.CategoryID = book.CategoryID
	bookDB.BookshelfID = book.BookshelfID
	bookDB.Publisher = book.Publisher
	fmt.Printf("BookDB OBJ :\n%v", bookDB)

	if err := db.Save(&bookDB).Error; err != nil {
		fmt.Printf("gagal mengupdate data book\n%v", err)
		return 0, err
	}

	return book.ID, nil
}

func DeleteBook(db *gorm.DB, book_id int) (uint, error) {
	//find by id firstt
	book := models.Book{}
	var err error

	if db.Find(&book, book_id).RecordNotFound() {
		fmt.Println("data book yang akan dihapus tidak ada.")
		return 0, nil
	}

	if err := db.Where("id = ? AND deleted_at IS NULL", book_id).Delete(models.Book{}).Error; err != nil {
		fmt.Println("gagal menghapus data book")
	}
	return book.ID, err
}
