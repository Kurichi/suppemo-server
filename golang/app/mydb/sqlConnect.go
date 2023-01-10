package mydb

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var db *sql.DB

func GetDB() *sql.DB {
	return db
}

func SqlConnect() error {
	err := godotenv.Load(fmt.Sprintf("%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		return err
	}
	DBMS := os.Getenv("DBMS")
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASS")
	HOST := os.Getenv("DB_HOST")
	PORT := os.Getenv("DB_PORT")
	DBNAME := os.Getenv("DB_NAME")

	CONNECT := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Asia%%2FTokyo",
		USER, PASS, HOST, PORT, DBNAME)

	if db, err = sql.Open(DBMS, CONNECT); err != nil {
		return err
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(10 * time.Second)

	if err := db.Ping(); err != nil {
		return err
	}

	fmt.Println("DB接続成功")

	return nil
}
