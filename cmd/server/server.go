package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func Server(logger *zap.Logger) error {
	r := chi.NewRouter()
	r.Mount("/", Router(logger))

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		return err
	} else {
		fmt.Println("Server is running on PORT: 3000")
	}

	return nil
}
