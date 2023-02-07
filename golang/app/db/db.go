package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func New() *sql.DB {
	return NewDB()
}

func NewDB() *sql.DB {
	err := godotenv.Load(fmt.Sprintf("%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Fatal(err)
	}
	DBMS := os.Getenv("DBMS")
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASS")
	HOST := os.Getenv("DB_HOST")
	PORT := os.Getenv("DB_PORT")
	DBNAME := os.Getenv("DB_NAME")

	CONNECT := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Asia%%2FTokyo",
		USER, PASS, HOST, PORT, DBNAME)

	db, err := sql.Open(DBMS, CONNECT)
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(10 * time.Second)

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}
