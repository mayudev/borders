package transport

import "github.com/melsincostan/borders/db/models"

var Default = []models.TransportBase{
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
