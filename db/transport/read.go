package transport

import (
	"github.com/melsincostan/borders/db/models"
	"gorm.io/gorm"
)

func Read(db *gorm.DB) (res []models.Transport, err error) {
	if err := db.Model(&models.Transport{}).Scan(&res).Error; err != nil {
		return nil, err
	}
	return
}
