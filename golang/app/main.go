package main

import (
	"suppemo-api/handler"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	e := handler.NewRouter()

	e.Logger.Fatal(e.Start(":8080"))
}
