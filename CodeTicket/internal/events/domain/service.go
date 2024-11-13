package domain

import (
	"errors"
	"fmt"
)

var (
	ErrSpotServiceQuantityZero = errors.New("quantity must be greater than zero")
)

type SpotService struct{}

// NewSpotService creates a new SpotService.
func NewSpotService() *SpotService {
	return &SpotService{}
}

// GenerateSpots generates the specified number of spots for an event.
func (s *SpotService) GenerateEventSpots(event *Event, quantity int) error {
	if quantity <= 0 {
		return ErrSpotServiceQuantityZero
	}

	for i := 0; i < quantity; i++ {
		spotName := generateSpotName(i)
		event.AddSpot(spotName)
	}

	return nil
}

func (s *SpotService) GenerateSpots(event *Event, index int) (*Spot, error) {
	spotName := generateSpotName(index)
	spot, err := NewSpot(event, spotName)
	if err != nil {
		return nil, err
	}

	return spot, nil
}

func generateSpotName(index int) string {
	// Gera um nome de spot baseado no índice. Ex: A1, A2, ..., B1, B2, etc.
	letter := 'A' + rune(index/10)
	number := index%10 + 1
	return fmt.Sprintf("%c%d", letter, number)
}
