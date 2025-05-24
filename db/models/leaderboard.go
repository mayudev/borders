package models

type LeaderboardEntry[T any] struct {
	Item         T `gorm:"embedded"`
	Ratio        float32
	Crossings    uint64
	Checks       uint64
	PapersChecks uint64
}
