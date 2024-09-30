package services

import (
	"errors"
	"pastebin-clone/configs"
	"pastebin-clone/internal/repositories"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceInterface interface {
	RegisterUser(username string, password string) (uuid.UUID, error)
	Login(username string, password string) (string, string, string, error)
	RefreshAccessToken(refreshToken string) (string, string, error)
}

type AuthService struct {
	Repo repositories.AuthRepositoryInterface // Dependency Injection ile repository alınıyor
}

func NewAuthService(repo repositories.AuthRepositoryInterface) AuthServiceInterface {
	return &AuthService{
		Repo: repo,
	}
}

func (s *AuthService) RegisterUser(username string, password string) (uuid.UUID, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return s.Repo.CreateUser(username, string(hashedPassword))
}

func (s *AuthService) Login(username string, password string) (string, string, string, error) {
	storedUser, err := s.Repo.GetUserByUsername(username)
	if err != nil {
		return "", "", "", errors.New("Invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(password)); err != nil {
		return "", "", "", errors.New("Invalid credentials")
	}

	accessToken, accessTokenClaims, err := createToken(storedUser.ID.String(), time.Now().Add(15*time.Minute).Unix(), "access")
	if err != nil {
		return "", "", "", err
	}
	refreshToken, _, err := createToken(storedUser.ID.String(), time.Now().Add(7*24*time.Hour).Unix(), "refresh")
	if err != nil {
		return "", "", "", err
	}

	expireDate := time.Unix(accessTokenClaims["exp"].(int64), 0).Format(time.RFC3339)

	return accessToken, refreshToken, expireDate, nil
}

func (s *AuthService) RefreshAccessToken(refreshToken string) (string, string, error) {
	claims, err := validateToken(refreshToken, "refresh")
	if err != nil {
		return "", "", err
	}

	userID := claims["sub"].(string)
	accessToken, accessTokenClaims, err := createToken(userID, time.Now().Add(15*time.Minute).Unix(), "access")
	if err != nil {
		return "", "", err
	}

	expireDate := time.Unix(accessTokenClaims["exp"].(int64), 0).Format(time.RFC3339)

	return accessToken, expireDate, nil
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
