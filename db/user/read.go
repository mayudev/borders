package user

import (
	"github.com/melsincostan/borders/db/models"
	"gorm.io/gorm"
)

func Authenticate(db *gorm.DB, username, password string) (e error) {
	var user models.User

	if err := db.Model(&models.User{}).Where("name = ?", username).First(&user).Error; err != nil {
		return err
	}
	return user.Check(password)
}
