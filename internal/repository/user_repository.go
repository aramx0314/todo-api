package repository

import (
	"database/sql"
	"todo-api/internal/model"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) FindByUsername(username string) (*model.User, error) {
	row := r.DB.QueryRow("SELECT username, password FROM users WHERE username=$1", username)
	var user model.User
	err := row.Scan(&user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Create(user model.User) error {
	_, err := r.DB.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", user.Username, user.Password)
	return err
}
