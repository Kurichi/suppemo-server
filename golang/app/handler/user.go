package handler

import (
	"fmt"
	"net/http"
	"suppemo-api/middleware"
	middlware "suppemo-api/middleware"

	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	_, err := middlware.Auth(authHeader)
	if err != nil {
		fmt.Printf("[ERROR] %v", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	uid := c.QueryParam("uid")

	user, err := middleware.GetUser(uid)
	if err != nil {
		fmt.Printf("[ERROR] %v", err.Error())
		return c.String(http.StatusBadRequest, "no-user")
	}

	type result struct {
		UID    string `json:"uid"`
		Name   string `json:"name"`
		Avatar string `json:"avatar"`
	}
	return c.JSON(http.StatusOK, result{
		UID:    user.UID,
		Name:   user.DisplayName,
		Avatar: user.PhotoURL,
	})
}
