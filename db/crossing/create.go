package crossing

import (
	"errors"

	"github.com/melsincostan/borders/db/models"
	"gorm.io/gorm"
)

var (
	ErrIdCheck = errors.New("a crossing with an ID check implies a border check")
)

func Create(db *gorm.DB, cb models.CrossingBase, countryID, transportID uint) (c *models.Crossing, e error) {
	if cb.PapersCheck && !cb.BorderCheck {
		return nil, ErrIdCheck
	}
	c = &models.Crossing{
		CrossingBase: cb,
		CountryID:    countryID,
		TransportID:  transportID,
	}

	if err := db.Model(&models.Crossing{}).Create(c).Error; err != nil {
		return nil, err
	}
	return
}
