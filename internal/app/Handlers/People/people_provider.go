package peopleHandler

import "go.uber.org/zap"

func InitPeopleHandler(logger *zap.Logger) PeopleHandler {
	handler := NewPoepleHandlerImpl(logger)
	return handler
}
