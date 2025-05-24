package transport

import (
	"github.com/melsincostan/borders/db/models"
	"gorm.io/gorm"
)

func Delete(db *gorm.DB, id uint) (e error) {
	if err := db.Model(&models.Transport{}).Where("id = ?", id).Delete(&models.Transport{}).Error; err != nil {
		return err
	}
	return
}
