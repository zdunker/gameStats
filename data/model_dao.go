package data

// readings

type ReadingDataHelper struct{}

var _ readingDataService = (*ReadingDataHelper)(nil)

func (h *ReadingDataHelper) GetReadings(pageNo int) ([]Reading, error) {
	return []Reading{}, nil
}

func (h *ReadingDataHelper) PostReadings(readings []Reading) error {
	return nil
}
