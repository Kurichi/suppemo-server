package service

import (
	"context"
	"fmt"
	"strings"
	"suppemo-api/model"
)

func (s *Service) CreateUser(ctx context.Context, uid string) (*model.User, error) {
	const (
		insert  = `INSERT INTO users(uid) VALUES(?)`
		confirm = `SELECT firebase_uid, created_at, updated_at FROM users WHERE id = ?`
	)

	conn, err := s.db.Conn(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, insert)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	r, err := stmt.ExecContext(ctx, uid)
	if err != nil {
		return nil, err
	}

	id, err := r.LastInsertId()
	if err != nil {
		return nil, err
	}

	row := conn.QueryRowContext(ctx, confirm, id)
	user := &model.User{}
	if err := row.Scan(user.FirebaseUID, user.CreatedAt, user.UpdatedAt); err != nil {
		return nil, err
	}
	user.ID = id

	return user, nil
}

func (s *Service) ReadUser(ctx context.Context, uid string) (*model.User, error) {
	const (
		read = `SELECT id, created_at, updated_at FROM users WHERE firebase_uid = ?`
	)

	conn, err := s.db.Conn(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, read)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, uid)
	user := &model.User{}
	if err := row.Scan(user.ID, user.CreatedAt, user.UpdatedAt); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) FindUser(ctx context.Context, uids []string) ([]*model.User, error) {
	const (
		findFmt = `SELECT firebase_uid, name, photo_url FROM users WHERE firebase_uid IN (?%s)`
	)
	conn, err := s.db.Conn(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, fmt.Sprintf(findFmt, strings.Repeat(",?", len(uids)-1)))
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	list := make([]interface{}, len(uids))
	for i, uid := range uids {
		list[i] = uid
	}

	rows, err := stmt.QueryContext(ctx, list)
	users := []*model.User{}
	for rows.Next() {
		user := &model.User{}
		rows.Scan(user.FirebaseUID, user.Name, user.PhotoURL)
		users = append(users, user)
	}

	return users, nil
}
