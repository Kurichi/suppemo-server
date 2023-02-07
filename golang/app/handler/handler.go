package handler

import (
	"context"
	"strings"
	"suppemo-api/service"

	"firebase.google.com/go/auth"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	service service.Service
	client  *auth.Client
}

func NewHandler(
	service *service.Service,
	client *auth.Client,
) *Handler {
	return &Handler{
		service: *service,
		client:  client,
	}
}

func (h *Handler) FirebaseAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		idToken := strings.Split(authHeader, "Bearer ")[1]
		auth, err := h.client.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			return err
		}
		h.service.UID = auth.UID

		return next(c)
	}
}
