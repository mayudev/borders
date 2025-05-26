package models

import (
	"time"

	"gorm.io/gorm"
)

type Crossing struct {
	gorm.Model
	CrossingBase
	CountryID   uint
	Country     Country
	TransportID uint
	Transport   Transport
	UserID      uint
	User        User
}

type CrossingBase struct {
	When        time.Time
	BorderCheck bool
	PapersCheck bool
}
