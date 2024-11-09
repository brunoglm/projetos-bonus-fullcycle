package domain

import (
	"errors"
	"fmt"
)

var (
	ErrSpotServiceQuantityZero = errors.New("quantity must be greater than zero")
)

type spotService struct{}

// NewSpotService creates a new SpotService.
func NewSpotService() *spotService {
	return &spotService{}
}

// GenerateSpots generates the specified number of spots for an event.
func (s *spotService) GenerateSpots(event *Event, quantity int) error {
	if quantity <= 0 {
		return ErrSpotServiceQuantityZero
	}

	for i := 0; i < quantity; i++ {
		spotName := fmt.Sprintf("%c%d", 'A'+i/10, i%10+1) // Generate spot name like A1, A2, ..., B1, B2, ...
		event.AddSpot(spotName)
	}

	return nil
}
