package repositories

import (
	"fmt"
	models "libraryapi/models"

	"github.com/jinzhu/gorm"
)

func FindAllBookshelf(db *gorm.DB) ([]models.Bookshelf, error) {
	bookshelves := []models.Bookshelf{}

	if err := db.Find(&bookshelves).Error; err != nil {
		return nil, err
	}

	return bookshelves, nil
}

func FindBookshelfByID(db *gorm.DB, id uint) ([]models.Bookshelf, error) {
	bookshelf := []models.Bookshelf{}

	if err := db.Find(&bookshelf, id).Error; err != nil {
		return nil, err
	}

	return bookshelf, nil
}

func InsertBookshelf(db *gorm.DB, bks models.Bookshelf) (uint, error) {
	if err := db.Create(&bks).Error; err != nil {
		fmt.Println("ada error di insert bookshelf :V \n%v", err)
		return 0, err
	}

	return bks.ID, nil
}

func UpdateBookshelf(db *gorm.DB, bookshelf *models.Bookshelf) (uint, error) {
	bookshelfDB := models.Bookshelf{}

	fmt.Println("Obj Bookshelf :", bookshelf)

	// if err := db.Where("id = ?", bookshelf.ID).First(bookshelf).Error; err != nil {
	// 	fmt.Println("User tidak ditemukan\ntidak ada bookshelf yang a m a n")
	// 	return 0, err
	// }

	if db.Where("id = ?", bookshelf.ID).First(&bookshelfDB).RecordNotFound() {
		fmt.Println("User tidak ditemukan\ntidak ada bookshelf yang a m a n")
		return 0, nil
	}

	//set data bookshelf dari parameter yang dikirim
	bookshelfDB.Code = bookshelf.Code
	fmt.Printf("BookshelfDB OBJ :\n%v", bookshelfDB)

	if err := db.Save(&bookshelfDB).Error; err != nil {
		fmt.Printf("gagal mengupdate data bookshelf\n%v", err)
		return 0, err
	}

	return bookshelf.ID, nil
}

func DeleteBookshelf(db *gorm.DB, bookshelf_id int) (uint, error) {
	//find by id firstt
	bookshelf := models.Bookshelf{}
	var err error

	if db.Find(&bookshelf, bookshelf_id).RecordNotFound() {
		fmt.Println("data bookshelf yang akan dihapus tidak ada.")
		return 0, nil
	}

	if err := db.Where("id = ? AND deleted_at IS NULL", bookshelf_id).Delete(models.Bookshelf{}).Error; err != nil {
		fmt.Println("gagal menghapus data bookshelf")
	}
	return bookshelf.ID, err
}
