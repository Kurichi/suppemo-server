package handler

import (
	"net/http"
	middleware "suppemo-api/middleware"
	"suppemo-api/model"

	"github.com/labstack/echo/v4"
)

func InitHandler(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	user, err := middleware.Auth(authHeader)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	type BodyType struct {
		PushToken string `json:"push_token"`
	}
	body := &BodyType{}
	if err = c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Create userif not exists
	if err = model.CreateUser(user.UID); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if body.PushToken == "" {
		return c.String(http.StatusOK, "not push")
	}

	if err = model.CreatePushToken(user.UID, body.PushToken); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, "push")
}
