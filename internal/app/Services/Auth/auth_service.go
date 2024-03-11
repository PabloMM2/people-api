package authService

import (
	"fmt"
	"os"
	models "people-api/internal/app/Models"
	authDto "people-api/internal/app/dtos/Auth"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	GetToken(req *authDto.AuthRequest, person *models.Person) (*authDto.TokenResponse, error)
}

type AuthServiceImpl struct {
	Logger *zap.Logger
}

func NewAuthServiceImpl(logger *zap.Logger) AuthService {
	return &AuthServiceImpl{Logger: logger}
}

func (ps *AuthServiceImpl) GetToken(req *authDto.AuthRequest, person *models.Person) (*authDto.TokenResponse, error) {
	err := bcrypt.CompareHashAndPassword([]byte(*person.Password), []byte(*req.Password))
	if err != nil {
		return nil, err
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = fmt.Sprintf("%s, %s", *person.Name, *person.LastName)
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}

	tokenResponse := &authDto.TokenResponse{Token: &tokenString}

	return tokenResponse, nil
}
