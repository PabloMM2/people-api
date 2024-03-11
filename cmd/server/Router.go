package server

import (
	"net/http"
	authHandler "people-api/internal/app/Handlers/Auth"
	peopleHandler "people-api/internal/app/Handlers/People"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func Router(logger *zap.Logger) *chi.Mux {

	// Init handlers with routers
	peopleHan := peopleHandler.InitPeopleHandler(logger)
	authHan := authHandler.InitAuthHandler(logger)

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World!"))
	})
	r.Mount("/people", peopleHan.Router())
	r.Mount("/auth", authHan.Router())
	return r
}
