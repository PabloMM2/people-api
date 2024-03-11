package authService

import "go.uber.org/zap"

func InitAuthService(logger *zap.Logger) AuthService {
	return &AuthServiceImpl{Logger: logger}
}
