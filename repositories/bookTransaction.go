package repositories

import (
	"fmt"
	models "libraryapi/models"
	"strconv"

	"github.com/jinzhu/gorm"
)

func FindAllBookTransaction(db *gorm.DB) ([]models.BookTransaction, error) {
	bookTransactions := []models.BookTransaction{}

	// if err := db.Find(&bookTransactions).Error; err != nil {
	if err := db.Preload("Book").Find(&bookTransactions).Error; err != nil {
		return nil, err
	}

	return bookTransactions, nil
}

func InsertBookTransaction(db *gorm.DB, bkTrx *models.BookTransaction) (uint, error) {
	if err := db.Create(&bkTrx).Error; err != nil {
		fmt.Println("ada error di insert book transaction :V \n%v", err)
		return 0, err
	}

	return bkTrx.ID, nil
}

func ReturnBook(db *gorm.DB, id uint) (uint, error) {
	//find by id firstt
	booktrx := models.BookTransaction{}
	var err error

	fmt.Println("id usr :", strconv.Itoa(int(id)))

	if db.Find(&booktrx, strconv.Itoa(int(id))).RecordNotFound() {
		fmt.Println("data buku yang ingin dikembalikan tidak ada.")
		return 0, nil
	}

	if err := db.Where("id=?", id).Model(&booktrx).Update("Status", "Dikembalikan"); err != nil {
		fmt.Println("gagal mengembalikan buku")
	}
	return id, err
}
