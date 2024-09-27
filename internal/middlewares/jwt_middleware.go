package middlewares

import (
	"net/http"
	"pastebin-clone/configs"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

// JWT doğrulama middleware
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

		return next(c)
	}
}
