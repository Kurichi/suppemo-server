package model

import (
	"fmt"
	"suppemo-api/mydb"
	"time"
)

type Friend struct {
	UID       string    `json:"uid" form:"uid" query:"uid"`
	FriendUID string    `json:"friend_uid" form:"friend_uid" query:"friend_uid"`
	Created   time.Time `json:"created" form:"created" query:"created"`
}

func CreateFriend(uid string, fuid string) error {
	db := mydb.GetDB()

	stmt, err := db.Prepare("INSERT INTO friends(uid, friend_uid) VALUES(?,?)")
	if err != nil {
		fmt.Printf("[ERROR] friend prepare error: %v", err)
		return err
	}

	if _, err = stmt.Exec(uid, fuid); err != nil {
		fmt.Printf("[ERROR] friend exec error: %v", err)
		return err
	}

	return nil
}

func FindFriends(uid string) ([]string, error) {
	db := mydb.GetDB()

	stmt, err := db.Prepare("SELECT friend_uid FROM friends WHERE uid = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var friend_uids []string
	rows, err := stmt.Query(uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var fuid string
		if err := rows.Scan(&fuid); err != nil {
			return nil, err
		}
		friend_uids = append(friend_uids, fuid)
	}

	stmt, err = db.Prepare("SELECT uid FROM friends WHERE friend_uid = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err = stmt.Query(uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var fuid string
		if err := rows.Scan(&fuid); err != nil {
			return nil, err
		}
		friend_uids = append(friend_uids, fuid)
	}

	return friend_uids, nil
}
