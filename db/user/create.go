package user

import (
	"fmt"

	"github.com/melsincostan/argon2id"
	"github.com/melsincostan/borders/db/models"
	"gorm.io/gorm"
)

func Create(db *gorm.DB, base models.UserBase, password string) (u *models.User, e error) {
	h, err := argon2id.New(password)
	if err != nil {
		return nil, fmt.Errorf("couldn't hash password: %w", err)
	}
	base.Password = h.Serialize()
	u = &models.User{
		UserBase: base,
	}
	if err := db.Model(&models.User{}).Create(u).Error; err != nil {
		return nil, err
	}
	return
}
