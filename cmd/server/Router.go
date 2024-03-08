package server

import (
	"net/http"
	peopleHandler "people-api/internal/app/Handlers/People"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func Router(logger *zap.Logger) *chi.Mux {

	// Init handlers with routers
	peopleHan := peopleHandler.InitPeopleHandler(logger)

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World!"))
	})
	r.Mount("/people", peopleHan.Router())
	return r
}
