package repositories

import (
	models "libraryapi/models"

	"github.com/jinzhu/gorm"
)

func FindAllGuestbook(db *gorm.DB) ([]models.Guestbook, error) {
	guestbooks := []models.Guestbook{}

	if err := db.Find(&guestbooks).Error; err != nil {
		return nil, err
	}

	return guestbooks, nil
}

func FindGuestbookByID(db *gorm.DB, id uint) ([]models.Guestbook, error) {
	guestbooks := []models.Guestbook{}

	if err := db.Find(&guestbooks, id).Error; err != nil {
		return nil, err
	}

	return guestbooks, nil
}

func InsertGuestbook(db *gorm.DB, gbk *models.Guestbook) error {
	if err := db.Create(&gbk).Error; err != nil {
		return err
	}

	return nil
}
