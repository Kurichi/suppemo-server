package model

import (
	"fmt"
	"suppemo-api/mydb"
)

func CreatePushToken(uid string, token string) error {
	db := mydb.GetDB()

	stmt, err := db.Prepare("INSERT IGNORE INTO push_tokens(uid,token) VALUES(?,?)")
	if err != nil {
		fmt.Printf("[ERROR] token prepare error: %v", err)
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(uid, token); err != nil {
		fmt.Printf("[ERROR] token exec error: %v", err)
		return err
	}

	return nil
}

func FindPushTokens(uid string) ([]string, error) {
	db := mydb.GetDB()

	stmt, err := db.Prepare("SELECT token FROM push_tokens WHERE uid = ?")
	if err != nil {
		fmt.Printf("[ERROR] token prepare error: %v", err)
		return nil, err
	}
	defer stmt.Close()

	var tokens []string
	rows, err := stmt.Query(uid)
	if err != nil {
		fmt.Printf("[ERROR] token query error: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var token string
		if err = rows.Scan(&token); err != nil {
			fmt.Printf("[ERROR] token scan error: %v", err)
			return tokens, err
		}
		tokens = append(tokens, token)
	}

	return tokens, nil
}
