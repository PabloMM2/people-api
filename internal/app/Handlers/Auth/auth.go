package authHandler

import (
	"net/http"
	authService "people-api/internal/app/Services/Auth"
	peopleService "people-api/internal/app/Services/People"
	authDto "people-api/internal/app/dtos/Auth"

	"people-api/internal/app/utils"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type AuthHandler interface {
	Router() *chi.Mux
	Auth(w http.ResponseWriter, r *http.Request)
}

type AuthHandlerImpl struct {
	Logger *zap.Logger
}

func NewAuthHandlerImpl(logger *zap.Logger) AuthHandler {
	return &AuthHandlerImpl{Logger: logger}
}

func (ah *AuthHandlerImpl) Router() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/", ah.Auth)
	return r
}

func (ah *AuthHandlerImpl) Auth(w http.ResponseWriter, r *http.Request) {
	LOG := "AuthHandler.Auth."
	ah.Logger.Info(LOG + "Start")

	body := &authDto.AuthRequest{}
	errVal := utils.Validate(&r.Body, body)
	if errVal != nil {
		errMsg := errVal.Error()
		ah.Logger.Error(LOG+"Validator.", zap.Any("Error:", errMsg))
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unexpected error"))
		return
	}

	// Get the person from db using the service
	peopleService := peopleService.InitPeopleService(ah.Logger)
	person, err := peopleService.GetPersonByEmail(body.Email)
	if err != nil {
		ah.Logger.Error(LOG+"Error.find.Person", zap.Any("Error:", err))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unexpected error!"))
		return
	}

	// Compare credencials
	authService := authService.InitAuthService(ah.Logger)
	token, err := authService.GetToken(body, person)
	if err != nil {
		ah.Logger.Error(LOG+"Error.find.Person", zap.Any("Error:", err))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Credentials ar not valid!"))
		return
	}

	utils.SuccessResponse(w, token, nil)
}
