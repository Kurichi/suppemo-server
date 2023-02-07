package middleware

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
)

var (
	client *auth.Client
)

func Init() error {
	opt := option.WithCredentialsFile("service-account-file.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return err
	}

	client, err = app.Auth(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func Auth(idToken string) (*auth.UserRecord, error) {
	token, err := client.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		return nil, err
	}

	return client.GetUser(context.Background(), token.UID)
}

func GetUser(uid string) (*auth.UserRecord, error) {
	return client.GetUser(context.Background(), uid)
}

func FirebaseAuthMiddleware(client *auth.Client) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			return next(c)
		}
	}
}
