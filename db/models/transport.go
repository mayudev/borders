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

var DefaultTransports = []TransportBase{
	{Name: "walking"},
	{Name: "bike"},
	{Name: "car"},
	{Name: "city bus"},
	{Name: "regional bus"},
	{Name: "long distance bus"},
	{Name: "regional train"},
	{Name: "intercity train"},
	{Name: "high-speed train"},
	{Name: "sleeper train"},
	{Name: "tramway"},
	{Name: "tram-train"},
}
