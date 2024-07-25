package main

import (
	"database/sql"
	"net/http"

	httpHandler "github.com/caiof27/imersao-full-cycle/golang-api/internal/events/infra/http"
	"github.com/caiof27/imersao-full-cycle/golang-api/internal/events/infra/repository"
	"github.com/caiof27/imersao-full-cycle/golang-api/internal/events/infra/service"
	"github.com/caiof27/imersao-full-cycle/golang-api/internal/events/usecase"
)

func main() {
	db, err := sql.Open("mysql", "")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	eventRepo, err := repository.NewMysqlEventRepository(db)
	if err != nil {
		panic(err)
	}

	partnerBaseURLs := map[int]string{
		1: "http://localhost:3000",
		2: "http://localhost:3001",
	}

	partnerFactory := service.NewPartnerFactory(partnerBaseURLs)

	listEventsUseCase := usecase.NewListEventsUseCase(eventRepo)
	getEventUseCase := usecase.NewGetEventUseCase(eventRepo)
	listSpotsUseCase := usecase.NewListSpotsUseCase(eventRepo)
	buyTicketsUseCase := usecase.NewBuyTicketsUseCase(eventRepo, partnerFactory)

	eventsHandler := httpHandler.NewEventHandler(
		listEventsUseCase,
		getEventUseCase,
		buyTicketsUseCase,
		listSpotsUseCase,
	)

	r := http.NewServeMux()
	r.HandleFunc("GET /events", eventsHandler.ListEvents)
	r.HandleFunc("GET /events/{eventID}", eventsHandler.GetEvent)
	r.HandleFunc("GET /events/{eventID}/spots", eventsHandler.ListSpots)
	r.HandleFunc("POST /checkout", eventsHandler.BuyTickets)

	http.ListenAndServe(":8080", r)
}
