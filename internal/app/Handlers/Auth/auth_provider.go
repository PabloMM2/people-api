package authHandler

import "go.uber.org/zap"

func InitAuthHandler(logger *zap.Logger) AuthHandler {
	handler := NewAuthHandlerImpl(logger)
	return handler
}
