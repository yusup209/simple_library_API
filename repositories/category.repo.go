package repositories

import (
	"fmt"
	models "libraryapi/models"

	"github.com/jinzhu/gorm"
)

func FindAllCategory(db *gorm.DB) ([]models.Category, error) {
	categories := []models.Category{}

	if err := db.Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

func FindCategoryByID(db *gorm.DB, id uint) ([]models.Category, error) {
	category := []models.Category{}

	if err := db.Find(&category, id).Error; err != nil {
		return nil, err
	}

	return category, nil
}

func InsertCategory(db *gorm.DB, cat *models.Category) (uint, error) {
	if err := db.Create(&cat).Error; err != nil {
		fmt.Println("ada error di insert category :v\n%v", err)
		return 0, err
	}

	return cat.ID, nil
}

func UpdateCategory(db *gorm.DB, category *models.Category) (uint, error) {
	categoryDB := models.Category{}

	fmt.Println("Obj Users :", category)

	// if err := db.Where("id = ?", category.ID).First(category).Error; err != nil {
	// 	fmt.Println("User tidak ditemukan\ntidak ada category yang a m a n")
	// 	return 0, err
	// }

	if db.Where("id = ?", category.ID).First(&categoryDB).RecordNotFound() {
		fmt.Println("User tidak ditemukan\ntidak ada category yang a m a n")
		return 0, nil
	}

	//set data user dari parameter yang dikirim
	categoryDB.Name = category.Name
	fmt.Printf("UserDB OBJ :\n%v", categoryDB)

	if err := db.Save(&categoryDB).Error; err != nil {
		fmt.Printf("gagal mengupdate data category\n%v", err)
		return 0, err
	}

	return category.ID, nil
}

func DeleteCategory(db *gorm.DB, category_id int) (uint, error) {
	//find by id firstt
	category := models.Category{}
	var err error

	if db.Find(&category, category_id).RecordNotFound() {
		fmt.Println("data kategori yang akan dihapus tidak ada.")
		return 0, nil
	}

	if err := db.Where("id = ? AND deleted_at IS NULL", category_id).Delete(models.Category{}).Error; err != nil {
		fmt.Println("gagal menghapus data kategori")
	}
	return category.ID, err
}
