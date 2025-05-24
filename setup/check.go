package setup

import (
	"errors"
	"fmt"

	"github.com/melsincostan/borders/db/models"
	"gorm.io/gorm"
)

func IsSetup(db *gorm.DB) (err error) {
	if err := db.AutoMigrate(&models.Setup{}); err != nil {
		return fmt.Errorf("couldn't ensure presence of the table: %w", err)
	}
	var setup models.Setup
	if err := db.Model(&models.Setup{}).First(&setup).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrNotSetup
		} else {
			return fmt.Errorf("could not check db: %w", err)
		}
	}
	if setup.Version != version {
		return ErrWrongVersion
	}
	return
}
