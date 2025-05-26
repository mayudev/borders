package user

import (
	"github.com/melsincostan/borders/db/models"
	"gorm.io/gorm"
)

func Authenticate(db *gorm.DB, username, password string) (u *models.User, e error) {
	var user models.User

	if err := db.Model(&models.User{}).Where("name = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, user.Check(password)
}
