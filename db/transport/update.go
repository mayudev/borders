package transport

import (
	"github.com/melsincostan/borders/db/models"
	"gorm.io/gorm"
)

func Update(db *gorm.DB, tb models.TransportBase) (e error) {
	t := models.Transport{
		TransportBase: tb,
	}
	if err := db.Model(&models.Transport{}).Updates(t).Error; err != nil {
		return err
	}
	return
}
