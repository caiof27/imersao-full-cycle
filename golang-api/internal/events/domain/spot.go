package domain

import (
	"errors"

	"github.com/google/uuid"
)

type SpotStatus string

var (
	ErrInvalidSpotNumber       = errors.New("invalid spot number")
	ErrSpotNotFound            = errors.New("spot not found")
	ErrSpotAlreadyReserved     = errors.New("spot already reserved")
	ErrSpotNameTwoCharacters   = errors.New("spot name must be at least 2 characters long")
	ErrSpotNameRequired        = errors.New("spot name is required")
	ErrSpotNameStartWithLetter = errors.New("spot name must start with a letter")
	ErrSpotNameEndWithNumber   = errors.New("spot name must end with a number")
)

const (
	SpotStatusAvailable SpotStatus = "Available"
	SpotStatusSold      SpotStatus = "Sold"
)

type Spot struct {
	ID       string
	EventID  string
	Name     string
	Status   SpotStatus
	TicketID string
}

func NewSpot(event *Event, name string) (*Spot, error) {
	spot := &Spot{
		ID:      uuid.New().String(),
		EventID: event.ID,
		Name:    name,
		Status:  SpotStatusAvailable,
	}

	if err := spot.Validate(); err != nil {
		return nil, err
	}
	return spot, nil
}

func (s *Spot) Validate() error {
	if len(s.Name) == 0 {
		return ErrSpotNameRequired
	}
	if len(s.Name) < 2 {
		return ErrSpotNameTwoCharacters
	}
	if s.Name[0] < 'A' || s.Name[0] > 'Z' {
		return ErrSpotNameStartWithLetter
	}
	if s.Name[1] < '0' || s.Name[1] > '9' {
		return ErrSpotNameEndWithNumber
	}
	return nil
}

func (s *Spot) Reserve(tickedID string) error {
	if s.Status == SpotStatusSold {
		return ErrSpotAlreadyReserved
	}
	s.Status = SpotStatusSold
	s.TicketID = tickedID
	return nil
}
