package usecase

type SpotDTO struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	EventID  string `json:"event_id"`
	Reserved bool   `json:"reserved"`
	Status   string `json:"status"`
	TicketID string `json:"ticket_id"`
}

type EventDTO struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Location     string  `json:"location"`
	Organization string  `json:"organization"`
	Rating       string  `json:"rating"`
	Date         string  `json:"date"`
	Capacity     int     `json:"capacity"`
	Price        float64 `json:"price"`
	PartnerID    int     `json:"partner_id"`
	ImageURL     string  `json:"image_url"`
}
