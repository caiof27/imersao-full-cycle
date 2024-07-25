package domain

import (
	"errors"

	"github.com/google/uuid"
)

var (
	errTicketPriceZero   = errors.New("Ticket price must be greater than zero")
	ErrInvalidTicketKind = errors.New("invalid ticket kind")
)

type TicketType string

const (
	TicketTypeFull TicketType = "Full"
	TicketTypeHalf TicketType = "Half"
)

func IsValidTicketKind(ticketKind TicketType) bool {
	return ticketKind == TicketTypeFull || ticketKind == TicketTypeHalf
}

type Ticket struct {
	ID         string
	EventID    string
	Spot       *Spot
	TicketType TicketType
	Price      float64
}

func IsValidTicketType(ticketType TicketType) bool {
	return ticketType == TicketTypeFull || ticketType == TicketTypeHalf
}

func (t *Ticket) CalculatePrice() {
	if t.TicketType == TicketTypeHalf {
		t.Price /= 2
	}
}

func (t *Ticket) Validate() error {
	if t.Price <= 0 {
		return errTicketPriceZero
	}
	return nil
}

func NewTicket(event *Event, spot *Spot, ticketType TicketType) (*Ticket, error) {
	if !IsValidTicketKind(ticketType) {
		return nil, ErrInvalidTicketKind
	}

	ticket := &Ticket{
		ID:         uuid.New().String(),
		EventID:    event.ID,
		Spot:       spot,
		TicketType: ticketType,
		Price:      event.Price,
	}
	ticket.CalculatePrice()
	if err := ticket.Validate(); err != nil {
		return nil, err
	}
	return ticket, nil
}
