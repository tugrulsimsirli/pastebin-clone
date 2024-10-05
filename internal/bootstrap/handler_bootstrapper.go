package bootstrap

import (
	"pastebin-clone/internal/http/handlers"
	"pastebin-clone/internal/http/middlewares"
	"pastebin-clone/internal/repositories"
	"pastebin-clone/internal/services"

	"github.com/labstack/echo/v4"
)

func RegisterHandlers(e *echo.Echo) {
	// Repositories
	authRepo := repositories.NewAuthRepository()
	userRepo := repositories.NewUserRepository()
	snippetRepo := repositories.NewSnippetRepository()

	// Services
	authService := services.NewAuthService(authRepo, userRepo)
	userService := services.NewUserService(userRepo)
	snippetService := services.NewSnippetService(snippetRepo)

	// Handlers
	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)
	snippetHandler := handlers.NewSnippetHandler(snippetService)

	apiV1Auth := e.Group("/api/v1/auth")
	apiV1Snippet := e.Group("/api/v1/snippet")
	apiV1User := e.Group("/api/v1/user")

	// Auth endpoints
	apiV1Auth.POST("/register", authHandler.Register)
	apiV1Auth.POST("/login", authHandler.Login)
	apiV1Auth.POST("/refresh-token", authHandler.RefreshToken)

	// User endpoints
	apiV1User.GET("", userHandler.GetUserDetail, middlewares.JWTMiddleware)

	// Snippet endpoints
	apiV1Snippet.GET("", snippetHandler.GetSnippetsOwn, middlewares.JWTMiddleware)
	apiV1Snippet.GET("/user/:userId", snippetHandler.GetSnippetsByUserID, middlewares.JWTMiddleware)
	apiV1Snippet.GET("/:id", snippetHandler.GetSnippet, middlewares.JWTMiddleware)
	apiV1Snippet.POST("", snippetHandler.CreateSnippet, middlewares.JWTMiddleware)
	apiV1Snippet.PATCH("/:id", snippetHandler.UpdateSnippet, middlewares.JWTMiddleware)
	apiV1Snippet.DELETE("/:id", snippetHandler.DeleteSnippet, middlewares.JWTMiddleware)
}
