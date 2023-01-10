package model

import (
	"fmt"
	"suppemo-api/mydb"
	"time"
)

type User struct {
	UID     string    `json:"uid" form:"uid" query:"uid"`
	Created time.Time `json:"created" form:"created" query:"created"`
	Updated time.Time `json:"updated" form:"updated" query:"updated"`
}

func CreateUser(uid string) error {
	db := mydb.GetDB()

	stmt, err := db.Prepare("INSERT IGNORE INTO users(uid) VALUES(?)")
	if err != nil {
		fmt.Printf("[ERROR] user prepare error: %v", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(uid)
	if err != nil {
		fmt.Printf("[ERROR] user exec error: %v", err)
		return err
	}

	return nil
}

func FindUser(uid string) (bool, error) {
	db := mydb.GetDB()

	stmt, err := db.Prepare("SELECT EXISTS (SELECT uid FROM users WHERE uid = ?)")
	if err != nil {
		fmt.Printf("[ERROR] user prepare error: %v", err)
		return false, err
	}
	defer stmt.Close()

	var exists bool
	if err = stmt.QueryRow(uid).Scan(&exists); err != nil {
		fmt.Printf("[ERROR] user query error: %v", err)
		return false, err
	}

	return exists, nil
}
