package handler

import (
	"fmt"
	"net/http"
	middlware "suppemo-api/middleware"
	"suppemo-api/model"

	"github.com/labstack/echo/v4"
)

type requestBody struct {
	UID string `json:"uid" query:"uid" form:"uid"`
}

func AddFriend(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	user, err := middlware.Auth(authHeader)
	if err != nil {
		fmt.Printf("[ERROR] %v", err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	reqBody := new(requestBody)
	if err := c.Bind(reqBody); err != nil {
		fmt.Printf("[ERROR] %v", err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if user.UID == reqBody.UID {
		return c.String(http.StatusBadRequest, "you can't add friend yourself")
	}

	if err := model.CreateFriend(user.UID, reqBody.UID); err != nil {
		fmt.Printf("[ERROR] %v", err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, "add friend complete")
}

func GetFriends(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	user, err := middlware.Auth(authHeader)
	if err != nil {
		fmt.Printf("[ERROR] %v", err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	friends, err := model.FindFriends(user.UID)
	if err != nil {
		fmt.Printf("[ERROR] %v", err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	type response struct {
		UID         string `json:"_id"`
		DisplayName string `json:"name"`
		PhotoURL    string `json:"avatar"`
	}

	res := make([]response, len(friends))
	for i, friend := range friends {
		user, err := middlware.GetUser(friend)
		if err != nil {
			fmt.Printf("[ERROR] %v", err)
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		res[i].UID = friend
		res[i].DisplayName = user.DisplayName
		res[i].PhotoURL = user.PhotoURL
	}

	return c.JSON(http.StatusOK, res)
}
