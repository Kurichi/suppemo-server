package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

type MyClaim struct {
	UserId     string
	RefreshJti string
	jwt.StandardClaims
}

func main() {
	secretBytes, err := ioutil.ReadFile("./../../ssh-keys/secret.key")
	if err != nil {
		log.Fatal(err)
		return
	}

	secretKey, err := jwt.ParseRSAPrivateKeyFromPEM(secretBytes)
	if err != nil {
		log.Fatal(err)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, MyClaim{
		UserId: "10o",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		}})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(tokenString)
}
