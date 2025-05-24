package crossing

import (
	"github.com/melsincostan/borders/db/models"
	"gorm.io/gorm"
)

func Delete(db *gorm.DB, id uint) (e error) {
	if err := db.Model(&models.Crossing{}).Where("id = ?", id).Delete(&models.Crossing{}).Error; err != nil {
		return err
	}
	return
}
