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
	snippetRepo := repositories.NewSnippetRepository() // Snippet repository

	// Services
	authService := services.NewAuthService(authRepo, userRepo)
	snippetService := services.NewSnippetService(snippetRepo) // Snippet service

	// Handlers
	authHandler := handlers.NewAuthHandler(authService)
	snippetHandler := handlers.NewSnippetHandler(snippetService) // Snippet handler

	apiV1 := e.Group("/api/v1")

	// Auth endpoints
	apiV1.POST("/register", authHandler.Register)
	apiV1.POST("/login", authHandler.Login)
	apiV1.POST("/refresh-token", authHandler.RefreshToken)

	// Snippet endpoints
	apiV1.GET("/snippet", snippetHandler.GetSnippets, middlewares.JWTMiddleware)
	apiV1.GET("/snippet/:id", snippetHandler.GetSnippet, middlewares.JWTMiddleware)
	apiV1.POST("/snippet", snippetHandler.CreateSnippet, middlewares.JWTMiddleware)
	apiV1.PATCH("/snippet/:id", snippetHandler.UpdateSnippet, middlewares.JWTMiddleware)
	apiV1.DELETE("/snippet/:id", snippetHandler.DeleteSnippet, middlewares.JWTMiddleware)
}
