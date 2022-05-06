package data

import "time"

type readingDataService interface {
	GetReadings(pageNo int) ([]Reading, error)
	PostReadings(reading []Reading) error
}

type Reading struct {
	ID uint64

	Name     string
	Author   string
	Overview string

	ICN  string
	Data time.Time
}
