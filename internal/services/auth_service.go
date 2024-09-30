package services

import (
	"errors"
	"pastebin-clone/configs"
	"pastebin-clone/internal/http/models"
	"pastebin-clone/internal/repositories"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceInterface interface {
	RegisterUser(username string, password string) (uuid.UUID, error)
	Login(username string, password string) (*models.LoginResponseModel, error)
	RefreshAccessToken(refreshToken string) (*models.RefreshTokenResponseModel, error)
}

type AuthService struct {
	AuthRepo repositories.AuthRepositoryInterface
	UserRepo repositories.UserRepositoryInterface
}

func NewAuthService(authRepo repositories.AuthRepositoryInterface, userRepo repositories.UserRepositoryInterface) AuthServiceInterface {
	return &AuthService{
		AuthRepo: authRepo,
		UserRepo: userRepo,
	}
}

func (s *AuthService) RegisterUser(username string, password string) (uuid.UUID, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return uuid.Nil, err
	}
	return s.AuthRepo.CreateUser(username, string(hashedPassword))
}

func (s *AuthService) Login(username string, password string) (*models.LoginResponseModel, error) {
	storedUser, err := s.UserRepo.GetUserByUsername(username)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	accessToken, accessTokenClaims, err := createToken(storedUser.ID.String(), time.Now().Add(15*time.Minute).Unix(), "access")
	if err != nil {
		return nil, err
	}
	refreshToken, _, err := createToken(storedUser.ID.String(), time.Now().Add(7*24*time.Hour).Unix(), "refresh")
	if err != nil {
		return nil, err
	}

	expireDate := time.Unix(accessTokenClaims["exp"].(int64), 0).Format(time.RFC3339)

	response := &models.LoginResponseModel{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpireDate:   expireDate,
	}

	return response, nil
}

func (s *AuthService) RefreshAccessToken(refreshToken string) (*models.RefreshTokenResponseModel, error) {
	claims, err := validateToken(refreshToken, "refresh")
	if err != nil {
		return nil, err
	}

	userID := claims["sub"].(string)
	accessToken, accessTokenClaims, err := createToken(userID, time.Now().Add(15*time.Minute).Unix(), "access")
	if err != nil {
		return nil, err
	}

	expireDate := time.Unix(accessTokenClaims["exp"].(int64), 0).Format(time.RFC3339)

	response := &models.RefreshTokenResponseModel{
		AccessToken: accessToken,
		ExpireDate:  expireDate,
	}

	return response, nil
}

func createToken(subject string, expirationTime int64, tokenType string) (string, jwt.MapClaims, error) {
	accessTokenClaims := jwt.MapClaims{
		"sub":  subject,
		"exp":  expirationTime,
		"type": tokenType,
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)

	accessTokenString, err := accessToken.SignedString([]byte(configs.AppConfig.JWTSecretKey))

	return accessTokenString, accessTokenClaims, err
}

func validateToken(tokenString string, tokenType string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token signature")
		}
		return []byte(configs.AppConfig.JWTSecretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid or expired token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	if claims["type"] != tokenType {
		return nil, errors.New("invalid token type")
	}

	return claims, nil
}
