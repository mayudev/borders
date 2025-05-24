package country

import (
	"github.com/melsincostan/borders/db/models"
	"gorm.io/gorm"
)

func Update(db *gorm.DB, id uint, cb models.CountryBase) (e error) {
	c := models.Country{
		CountryBase: cb,
	}
	if err := db.Model(&models.Country{}).Updates(c).Error; err != nil {
		return err
	}
	return
}
