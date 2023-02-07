package model

import (
	"time"
)

type (
	User struct {
		ID          int64     `json:"id"`
		FirebaseUID string    `json:"firebase_uid"`
		Name        string    `json:"name"`
		PhotoURL    string    `json:"photo_url"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
		IsDelete    bool      `json:"is_delete"`
	}

	SignInRequest struct {
		FirebaseUID string `json:"firebase_uid"`
	}
	SignInResponse struct {
		User User `json:"user"`
	}

	SignUpRequest struct {
		FirebaseUID string `json:"firebase_uid"`
	}
	SignUpResponse struct {
		User User `json:"user"`
	}

	FindUserRequest struct {
		FirebaseUIDs []string `json:"firebase_uids"`
	}
	FindUserResponse struct {
		Users []*User `json:"users"`
	}

	DeleteUserRequest struct {
		FirebaseUID string `json:"firebase_uid"`
	}
	DeleteUserResponse struct {
		User User `json:"user"`
	}

	UnDeleteUserRequest struct {
		FirebaseUID string `json:"firebase_uid"`
	}
	UnDeleteUserResponse struct {
		User User `json:"user"`
	}
)
