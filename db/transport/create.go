package transport

import (
	"github.com/melsincostan/borders/db/models"
	"gorm.io/gorm"
)

func Create(db *gorm.DB, tb models.TransportBase) (t *models.Transport, e error) {
	t = &models.Transport{
		TransportBase: tb,
	}
	if err := db.Model(&models.Transport{}).Create(t).Error; err != nil {
		return nil, err
	}
	return
}

func CreateBatch(db *gorm.DB, tbs []models.TransportBase) (t []models.Transport, e error) {
	t = []models.Transport{}
	for _, tb := range tbs {
		t = append(t, models.Transport{TransportBase: tb})
	}
	if err := db.Model(&models.Transport{}).Create(&t).Error; err != nil {
		return nil, err
	}
	return
}
