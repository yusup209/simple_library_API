package repositories

import (
	"fmt"
	models "libraryapi/models"
	"strconv"

	"github.com/jinzhu/gorm"
)

func FindAllUsers(db *gorm.DB) ([]models.Users, error) {
	users := []models.Users{}

	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func FindUsersByID(db *gorm.DB, id uint) ([]models.Users, error) {
	user := []models.Users{}

	if err := db.Find(&user, id).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func InsertUser(db *gorm.DB, gbk *models.Users) (uint, error) {
	if err := db.Create(&gbk).Error; err != nil {
		fmt.Println("ada error di insert user :V \n%v", err)
		return 0, err
	}

	return gbk.ID, nil
}

func UpdateUser(db *gorm.DB, user *models.Users) (uint, error) {
	userDB := models.Users{}

	fmt.Println("Obj Users :", user)

	// if err := db.Where("id = ?", user.ID).First(user).Error; err != nil {
	// 	fmt.Println("User tidak ditemukan\ntidak ada user yang a m a n")
	// 	return 0, err
	// }

	if db.Where("id = ?", user.ID).First(&userDB).RecordNotFound() {
		fmt.Println("User tidak ditemukan\ntidak ada user yang a m a n")
		return 0, nil
	}

	//set data user dari parameter yang dikirim
	userDB.Email = user.Email
	userDB.Username = user.Username
	userDB.Password = user.Password
	userDB.Avatar = user.Avatar
	userDB.Role = user.Role
	fmt.Printf("UserDB OBJ :\n%v", userDB)

	if err := db.Save(&userDB).Error; err != nil {
		fmt.Printf("gagal mengupdate data user\n%v", err)
		return 0, err
	}

	return user.ID, nil
}

func DeleteUsers(db *gorm.DB, user_id int) (uint, error) {
	//find by id firstt
	user := models.Users{}
	var err error

	fmt.Println("id usr :", strconv.Itoa(user_id))
	if db.Find(&user, strconv.Itoa(user_id)).RecordNotFound() {
		fmt.Println("data user yang akan dihapus tidak ada.")
		return 0, nil
	}

	if err := db.Where("id = ? AND deleted_at IS NULL", strconv.Itoa(user_id)).Delete(models.Users{}).Error; err != nil {
		fmt.Println("gagal menghapus data user")
	}
	return user.ID, err
}
