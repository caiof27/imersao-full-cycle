package http

import (
	"encoding/json"
	"net/http"

	"github.com/caiof27/imersao-full-cycle/golang-api/internal/events/usecase"
)

type EventsHandler struct {
	listEventsUseCase *usecase.ListEventsUseCase
	getEventUsecase   *usecase.GetEventUseCase
	buyTicketUseCase  *usecase.BuyTicketsUseCase
	listSpotsUsecase  *usecase.ListSpotsUseCase
}

func NewEventHandler(
	listEventsUseCase *usecase.ListEventsUseCase,
	getEventUsecase *usecase.GetEventUseCase,
	buyTicketUseCase *usecase.BuyTicketsUseCase,
	listSpotsUsecase *usecase.ListSpotsUseCase,
) *EventsHandler {
	return &EventsHandler{
		listEventsUseCase: listEventsUseCase,
		getEventUsecase:   getEventUsecase,
		buyTicketUseCase:  buyTicketUseCase,
		listSpotsUsecase:  listSpotsUsecase,
	}
}

func (h *EventsHandler) ListEvents(w http.ResponseWriter, r *http.Request) {
	output, err := h.listEventsUseCase.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func (h *EventsHandler) GetEvent(w http.ResponseWriter, r *http.Request) {

	eventID := r.PathValue("eventID")
	input := usecase.GetEventInputDTO{ID: eventID}

	output, err := h.getEventUsecase.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)

}

func (h *EventsHandler) ListSpots(w http.ResponseWriter, r *http.Request) {

	eventID := r.PathValue("eventID")
	input := usecase.ListSpotsInputDTO{EventID: eventID}

	output, err := h.listSpotsUsecase.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)

}

func (h *EventsHandler) BuyTickets(w http.ResponseWriter, r *http.Request) {
	var input usecase.BuyTicketsInputDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	output, err := h.buyTicketUseCase.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}
