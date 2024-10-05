package handlers

import (
	"net/http"
	"pastebin-clone/internal/http/models"
	"pastebin-clone/internal/services"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserService services.UserServiceInterface
}

func NewUserHandler(userService services.UserServiceInterface) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

// GetUser godoc
// @Summary      Get user detail
// @Description  Retrieves all user data for the authenticated user
// @Tags         User
// @Accept       json
// @Produce      json
// @Success      200  {object} models.UserDetailResponseModel
// @Success      204  "No Content"
// @Failure      400  {object} models.ErrorResponse
// @Failure      500  {object} models.ErrorResponse
// @Router       /api/v1/user [get]
func (h *UserHandler) GetUserDetail(c echo.Context) error {
	userID := c.Get("userID").(uuid.UUID)
	user, err := h.UserService.GetUserDetail(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
	}
	if user == nil {
		return c.NoContent(http.StatusNoContent)
	}
	return c.JSON(http.StatusOK, user)
}
