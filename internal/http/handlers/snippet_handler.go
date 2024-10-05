// internal/http/handlers/snippet_handler.go
package handlers

import (
	"net/http"
	"pastebin-clone/internal/http/models"
	"pastebin-clone/internal/services"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type SnippetHandler struct {
	SnippetService services.SnippetServiceInterface
}

func NewSnippetHandler(snippetService services.SnippetServiceInterface) *SnippetHandler {
	return &SnippetHandler{
		SnippetService: snippetService,
	}
}

// GetSnippets godoc
// @Summary      Get user snippets
// @Description  Retrieves all snippets for the authenticated user
// @Tags         Snippet
// @Accept       json
// @Produce      json
// @Success      200  {object} []models.SnippetResponseModel
// @Failure      400  {object} models.ErrorResponse
// @Failure      500  {object} models.ErrorResponse
// @Router       /api/v1/snippet [get]
func (h *SnippetHandler) GetSnippetsOwn(c echo.Context) error {
	userID := c.Get("userID").(uuid.UUID)
	snippets, err := h.SnippetService.GetAllSnippetsOwn(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
	}
	if snippets == nil {
		return c.NoContent(http.StatusNoContent)
	}
	return c.JSON(http.StatusOK, snippets)
}

// GetSnippets godoc
// @Summary      Get user snippets
// @Description  Retrieves all snippets for the authenticated user
// @Tags         Snippet
// @Accept       json
// @Produce      json
// @Param        userId path string true "Snippet ID"
// @Success      200  {object} []models.SnippetResponseModel
// @Success      204  "No Content"
// @Failure      400  {object} models.ErrorResponse
// @Failure      500  {object} models.ErrorResponse
// @Router       /api/v1/snippet/user/{userId} [get]
func (h *SnippetHandler) GetSnippetsByUserID(c echo.Context) error {
	userID := uuid.MustParse(c.Param("userId"))
	snippets, err := h.SnippetService.GetAllSnippetsByUserID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
	}
	if snippets == nil {
		return c.NoContent(http.StatusNoContent)
	}
	return c.JSON(http.StatusOK, snippets)
}

// GetSnippet godoc
// @Summary      Get snippet by ID
// @Description  Retrieves a snippet for the authenticated user by ID
// @Tags         Snippet
// @Accept       json
// @Produce      json
// @Param        id path string true "Snippet ID"
// @Success      200  {object} models.SnippetResponseModel
// @Failure      400  {object} models.ErrorResponse
// @Failure      500  {object} models.ErrorResponse
// @Router       /api/v1/snippet/{id} [get]
func (h *SnippetHandler) GetSnippet(c echo.Context) error {
	userID := c.Get("userID").(uuid.UUID)
	snippetID := uuid.MustParse(c.Param("id"))
	snippet, err := h.SnippetService.GetSnippetByID(userID, snippetID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
	}
	if snippet == nil {
		return c.NoContent(http.StatusNoContent)
	}
	return c.JSON(http.StatusOK, snippet)
}

// CreateSnippet godoc
// @Summary      Create a new snippet
// @Description  Creates a new snippet for the authenticated user
// @Tags         Snippet
// @Accept       json
// @Produce      json
// @Param        snippet body models.CreateSnippetRequestModel true "Snippet data"
// @Success      201  {object} models.IdResponseModel
// @Failure      400  {object} models.ErrorResponse
// @Failure      500  {object} models.ErrorResponse
// @Router       /api/v1/snippet [post]
func (h *SnippetHandler) CreateSnippet(c echo.Context) error {
	userID := c.Get("userID").(uuid.UUID)
	var req models.CreateSnippetRequestModel
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "Invalid input"})
	}

	snippet, err := h.SnippetService.CreateSnippet(userID, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, snippet)
}

// UpdateSnippet godoc
// @Summary      Update a snippet
// @Description  Updates a snippet for the authenticated user by ID
// @Tags         Snippet
// @Accept       json
// @Produce      json
// @Param        id path string true "Snippet ID"
// @Param        snippet body models.UpdateSnippetRequestModel true "Snippet data"
// @Success      200  {object} models.SnippetResponseModel
// @Failure      400  {object} models.ErrorResponse
// @Failure      500  {object} models.ErrorResponse
// @Router       /api/v1/snippet/{id} [patch]
func (h *SnippetHandler) UpdateSnippet(c echo.Context) error {
	userID := c.Get("userID").(uuid.UUID)
	snippetID := uuid.MustParse(c.Param("id"))
	var req models.UpdateSnippetRequestModel
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "Invalid input"})
	}

	snippet, err := h.SnippetService.UpdateSnippet(userID, snippetID, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, snippet)
}

// UpdateSnippet godoc
// @Summary      Update a snippet
// @Description  Updates a snippet for the authenticated user by ID
// @Tags         Snippet
// @Accept       json
// @Produce      json
// @Param        id path string true "Snippet ID"
// @Param        snippet_is_public body models.BooleanRequestModel true "Snippet IsPublic data"
// @Success      200  {object} models.SnippetResponseModel
// @Failure      400  {object} models.ErrorResponse
// @Failure      500  {object} models.ErrorResponse
// @Router       /api/v1/snippet/{id} [patch]
func (h *SnippetHandler) UpdateSnippetIsPublic(c echo.Context) error {
	userID := c.Get("userID").(uuid.UUID)
	snippetID := uuid.MustParse(c.Param("id"))
	var req models.BooleanRequestModel
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "Invalid input"})
	}

	snippet, err := h.SnippetService.UpdateSnippetIsPublic(userID, snippetID, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, snippet)
}

// DeleteSnippet godoc
// @Summary      Delete a snippet
// @Description  Deletes a snippet for the authenticated user by ID
// @Tags         Snippet
// @Accept       json
// @Produce      json
// @Param        id path string true "Snippet ID"
// @Success      204  {object} nil
// @Failure      400  {object} models.ErrorResponse
// @Failure      500  {object} models.ErrorResponse
// @Router       /api/v1/snippet/{id} [delete]
func (h *SnippetHandler) DeleteSnippet(c echo.Context) error {
	userID := c.Get("userID").(uuid.UUID)
	snippetID := uuid.MustParse(c.Param("id"))

	err := h.SnippetService.DeleteSnippet(userID, snippetID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}
