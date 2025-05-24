package models

import "gorm.io/gorm"

type Transport struct {
	gorm.Model
	TransportBase
	Crossings []Crossing
}

type TransportBase struct {
	Name string
}
