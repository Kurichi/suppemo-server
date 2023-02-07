package handler

import (
	"net/http"
	"suppemo-api/model"

	"github.com/labstack/echo/v4"
)

func (h *Handler) SignUp(c echo.Context) error {
	req := &model.SignUpRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	user, err := h.service.CreateUser(c.Request().Context(), req.FirebaseUID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, &model.SignUpResponse{
		User: *user,
	})
}

func (h *Handler) SignIn(c echo.Context) error {
	req := &model.SignInRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	user, err := h.service.ReadUser(c.Request().Context(), req.FirebaseUID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	return c.JSON(http.StatusOK, &model.SignInResponse{
		User: *user,
	})
}

func (h *Handler) GetUser(c echo.Context) error {
	req := &model.FindUserRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	users, err := h.service.FindUser(c.Request().Context(), req.FirebaseUIDs)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	return c.JSON(http.StatusOK, &model.FindUserResponse{
		Users: users,
	})
}
