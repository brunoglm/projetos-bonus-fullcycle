package usecase

import "codeticket/internal/events/domain"

type CreateSpotsInputDTO struct {
	EventID       string `json:"event_id"`
	NumberOfSpots int    `json:"number_of_spots"`
}

type CreateSpotsOutputDTO struct {
	Spots []SpotDTO `json:"spots"`
}

type CreateSpotsUseCase struct {
	repo    domain.EventRepository
	service *domain.SpotService
}

func NewCreateSpotsUseCase(repo domain.EventRepository) *CreateSpotsUseCase {
	return &CreateSpotsUseCase{repo: repo}
}

func (uc *CreateSpotsUseCase) Execute(input CreateSpotsInputDTO) (*CreateSpotsOutputDTO, error) {
	event, err := uc.repo.FindEventByID(input.EventID)
	if err != nil {
		return nil, err
	}

	spots := make([]domain.Spot, input.NumberOfSpots)
	for i := 0; i < input.NumberOfSpots; i++ {
		spot, err := uc.service.GenerateSpots(event, i)
		if err != nil {
			return nil, err
		}
		if err := uc.repo.CreateSpot(spot); err != nil {
			return nil, err
		}
		spots[i] = *spot
	}

	spotDTOs := make([]SpotDTO, len(spots))
	for i, spot := range spots {
		spotDTOs[i] = SpotDTO{
			ID:       spot.ID,
			Name:     spot.Name,
			Status:   string(spot.Status),
			TicketID: spot.TicketID,
		}
	}

	return &CreateSpotsOutputDTO{Spots: spotDTOs}, nil
}
