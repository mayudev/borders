package country

import (
	"github.com/melsincostan/borders/db/models"
	"gorm.io/gorm"
)

func Read(db *gorm.DB) (res []models.Country, err error) {
	res = []models.Country{}
	if err := db.Model(&models.Country{}).Scan(&res).Error; err != nil {
		return nil, err
	}
	return
}
