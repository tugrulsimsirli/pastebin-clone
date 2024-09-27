package middlewares

import (
	"net/http"
	"pastebin-clone/configs"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenStr := c.Request().Header.Get("Authorization")

		if tokenStr == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Missing token"})
		}

		tokenStr = tokenStr[len("Bearer "):]

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
			}
			return []byte(configs.AppConfig.JWTSecretKey), nil
		})

		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid or expired token"})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid token claims"})
		}

		userIDStr, ok := claims["sub"].(string)
		if !ok {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid user ID in token"})
		}

		userID, err := uuid.Parse(userIDStr)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid user ID format"})
		}

		c.Set("userID", userID)

		return next(c)
	}
}
