package auth

import (
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	Login(username, password string) (int64, error)
	Register(username, password string) (int64, error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Login(username, password string) (int64, error) {
	var id int64
	err := r.db.QueryRowx("SELECT id FROM users WHERE username = ? AND password = ?", username, password).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *repository) Register(username, password string) (int64, error) {
	res, err := r.db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, password)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, nil
	}
	return id, nil
}
