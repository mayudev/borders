package crossing

import (
	"fmt"

	"github.com/melsincostan/borders/db/models"
	"gorm.io/gorm"
)

type Stats struct {
	TotalCrossings    int64
	TotalBorderChecks int64
	TotalIdChecks     int64
	WorstCountry      Worst[models.Country]
	WorstTransport    Worst[models.Transport]
}

type Worst[T any] struct {
	Item  T `gorm:"embedded"`
	Ratio float32
}

const (
	WorstTransportQuery = "SELECT transports.*, COUNT(DISTINCT(tbc.id)) / COUNT(DISTINCT(tc.id)) AS ratio FROM transports LEFT JOIN crossings AS tc ON tc.transport_id = transports.id LEFT JOIN crossings AS tbc ON tbc.transport_id = transports.id AND tbc.border_check = 1 GROUP BY transports.id ORDER BY ratio DESC LIMIT 1"
	WorstCountryQuery   = "SELECT countries.*, COUNT(DISTINCT(tbc.id)) / COUNT(DISTINCT(tc.id)) AS ratio FROM countries LEFT JOIN crossings AS tc ON tc.country_id = countries.id LEFT JOIN crossings AS tbc ON tbc.country_id = countries.id AND tbc.border_check = 1 GROUP BY countries.id ORDER BY ratio DESC LIMIT 1"
)

func GetStats(db *gorm.DB) (s *Stats, e error) {
	var totalCrossings int64
	var totalBorderChecks int64
	var totalIdChecks int64
	var worstCountry Worst[models.Country]
	var worstTransport Worst[models.Transport]

	if err := db.Model(&models.Crossing{}).Count(&totalCrossings).Error; err != nil {
		return nil, fmt.Errorf("couldn't get total crossings count: %w", err)
	}

	if err := db.Model(&models.Crossing{}).Where("border_check = ?", true).Count(&totalBorderChecks).Error; err != nil {
		return nil, fmt.Errorf("couldn't get total crossings with a border check count: %w", err)
	}

	if err := db.Model(&models.Crossing{}).Where("papers_check = ?", true).Count(&totalIdChecks).Error; err != nil {
		return nil, fmt.Errorf("couldn't get total crossings with papers check count: %w", err)
	}

	if err := db.Raw(WorstTransportQuery).Scan(&worstTransport).Error; err != nil {
		return nil, fmt.Errorf("couldn't get worst transport: %w", err)
	}

	if err := db.Raw(WorstCountryQuery).Scan(&worstCountry).Error; err != nil {
		return nil, fmt.Errorf("couldn't get worst country: %w", err)
	}

	s = &Stats{
		TotalCrossings:    totalCrossings,
		TotalBorderChecks: totalBorderChecks,
		TotalIdChecks:     totalIdChecks,
		WorstCountry:      worstCountry,
		WorstTransport:    worstTransport,
	}
	return
}
