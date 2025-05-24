package country

import (
	"github.com/melsincostan/borders/db/models"
	"gorm.io/gorm"
)

func Delete(db *gorm.DB, id uint) (e error) {
	if err := db.Model(&models.Country{}).Where("id = ?", id).Delete(&models.Country{}).Error; err != nil {
		return err
	}
	return
}
