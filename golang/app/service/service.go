package service

import "database/sql"

type Service struct {
	db  *sql.DB
	UID string
}

func NewService(
	db *sql.DB,
) *Service {
	return &Service{
		db: db,
	}
}
