package user

import (
	"errors"
	"fmt"

	"github.com/melsincostan/argon2id"
	"github.com/melsincostan/borders/db/models"
	"gorm.io/gorm"
)

var (
	ErrWrongPassword = errors.New("old password does not match")
)

func ChangePassword(db *gorm.DB, uid uint, oldpw, newpw string) (err error) {
	var usr models.User

	if err := db.Model(&models.User{}).Where("id = ?", uid).First(&usr).Error; err != nil {
		return err
	}

	if err := usr.Check(oldpw); err != nil {
		return fmt.Errorf("%w: %w", ErrWrongPassword, err)
	}

	newpwh, err := argon2id.New(newpw)
	if err != nil {
		return fmt.Errorf("could not hash new password: %w", err)
	}

	if err := db.Model(&models.User{}).Where("id = ?", uid).UpdateColumn("password", newpwh.Serialize()).Error; err != nil {
		return err
	}
	return
}
