package country

import (
	"github.com/melsincostan/borders/db/models"
	"gorm.io/gorm"
)

func Create(db *gorm.DB, cb models.CountryBase) (c *models.Country, e error) {
	c = new(models.Country)
	c.CountryBase = cb
	if err := db.Model(&models.Country{}).Create(c).Error; err != nil {
		return nil, err
	}
	return
}

func CreateBatch(db *gorm.DB, cbs []models.CountryBase) (c []models.Country, e error) {
	c = []models.Country{}
	for _, cb := range cbs {
		c = append(c, models.Country{
			CountryBase: cb,
		})
	}
	if err := db.Model(&models.Country{}).Create(&c).Error; err != nil {
		return nil, err
	}
	return
}
