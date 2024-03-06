package server

import (
	peopleHandler "people-api/internal/app/Handlers/People"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func Router(logger *zap.Logger) *chi.Mux {

	// Init handlers with routers
	peopleHan := peopleHandler.InitPeopleHandler(logger)

	r := chi.NewRouter()
	r.Mount("/people", peopleHan.Router())
	return r
}
