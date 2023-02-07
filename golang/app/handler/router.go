package handler

import (
	"context"
	"suppemo-api/db"
	"suppemo-api/service"

	firebase "firebase.google.com/go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/api/option"
)

func NewRouter() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Firebase Auth Client 初期化
	opt := option.WithCredentialsFile("service-account-file.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic(err)
	}
	client, err := app.Auth(context.Background())
	if err != nil {
		panic(err)
	}

	handler := NewHandler(service.NewService(db.New()), client)

	api := e.Group("/api", handler.FirebaseAuthMiddleware)
	api.POST("/signup", handler.SignUp)
	api.GET("/signin", handler.SignIn)
	api.GET("/users", handler.GetUser)

	return e
}
