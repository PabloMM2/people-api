package peopleHandler

import (
	"net/http"
	peopleService "people-api/internal/app/Services/People"
	peopleDto "people-api/internal/app/dtos/People"
	"strconv"

	"people-api/internal/app/utils"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type PeopleHandler interface {
	Router() *chi.Mux
	GetPeople(w http.ResponseWriter, r *http.Request)
	CreatePeople(w http.ResponseWriter, r *http.Request)
	UpdatePeople(w http.ResponseWriter, r *http.Request)
	DeletePeople(w http.ResponseWriter, r *http.Request)
}

type PoepleHandlerImpl struct {
	Logger *zap.Logger
}

func NewPoepleHandlerImpl(logger *zap.Logger) PeopleHandler {
	return &PoepleHandlerImpl{Logger: logger}
}

func (ph *PoepleHandlerImpl) Router() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/{id}", ph.GetPeople)
	r.Post("/", ph.CreatePeople)
	r.Put("/", ph.UpdatePeople)
	r.Delete("/", ph.DeletePeople)
	return r
}

func (ph *PoepleHandlerImpl) GetPeople(w http.ResponseWriter, r *http.Request) {
	LOG := "PeopleHandler.GetPeople."
	ph.Logger.Info(LOG + "Start")

	peopleId := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(peopleId, 10, 32)
	if err != nil {
		ph.Logger.Error(LOG+"Error.Parse.string.to.uint.", zap.Any("Error:", err))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Id error to find person"))
		return
	}

	service := peopleService.InitPeopleService(ph.Logger)
	personId := uint(id)
	person, err := service.GetPerson(&personId)
	if err != nil {
		ph.Logger.Error(LOG+"Error.find.Person", zap.Any("Error:", err))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error to find person"))
		return
	}

	utils.SuccessResponse(w, person, nil)
}

func (ph *PoepleHandlerImpl) CreatePeople(w http.ResponseWriter, r *http.Request) {
	LOG := "PeopleHandler.CreatePeople."
	ph.Logger.Info(LOG + "Start")
	// validate the dro request
	body := &peopleDto.PeopleCreateRquest{}
	errVal := utils.Validate(&r.Body, body)
	if errVal != nil {
		errMsg := errVal.Error()
		ph.Logger.Error(LOG+"Validator.", zap.Any("Error:", errMsg))
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error in Validator"))
		return
	}

	ph.Logger.Info(LOG+"body", zap.Any("Data:", body))

	service := peopleService.InitPeopleService(ph.Logger)
	person, err := service.CreatePerson(body)

	if err != nil {
		ph.Logger.Error(LOG+"Validator.", zap.Any("Error:", err))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error Creatin person"))
		return
	}

	ph.Logger.Info(LOG+"PersonCreated", zap.Any("Person:", person))
	statusCreated := http.StatusCreated
	utils.SuccessResponse(w, person, &statusCreated)
}

func (ph *PoepleHandlerImpl) UpdatePeople(w http.ResponseWriter, r *http.Request) {
	//TODO
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Hello World from people!"))
}

func (ph *PoepleHandlerImpl) DeletePeople(w http.ResponseWriter, r *http.Request) {
	//TODO
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Hello World from people!"))
}
