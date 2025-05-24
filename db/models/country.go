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

var DefaultCountries = []CountryBase{
	{Name: "Germany"},
	{Name: "France"},
	{Name: "Poland"},
	{Name: "Austria"},
	{Name: "Switzerland"},
}
