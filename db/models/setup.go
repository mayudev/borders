package models

import (
	"gorm.io/gorm"
)

type Setup struct {
	gorm.Model
	Version uint
}
