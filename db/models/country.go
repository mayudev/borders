package models

import "gorm.io/gorm"

type Country struct {
	gorm.Model
	CountryBase
	Crossings []Crossing
}

type CountryBase struct {
	Name string
}
