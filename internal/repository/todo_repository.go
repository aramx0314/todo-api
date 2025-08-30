package repository

import (
	"database/sql"
	"errors"
	"todo-api/internal/model"
)

type TodoRepository struct {
	DB *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{DB: db}
}

func (r *TodoRepository) Create(todo model.Todo) error {
	_, err := r.DB.Exec("INSERT INTO todos (id, title, completed, username) VALUES ($1, $2, $3, $4)",
		todo.Id, todo.Title, todo.Completed, todo.Username)
	return err
}

func (r *TodoRepository) FindByUsername(username string) ([]model.Todo, error) {
	rows, err := r.DB.Query("SELECT id, title, completed, username FROM todos WHERE username=$1", username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := []model.Todo{}
	for rows.Next() {
		var t model.Todo
		if err := rows.Scan(&t.Id, &t.Title, &t.Completed, &t.Username); err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}
	return todos, nil
}

func (r *TodoRepository) Update(todo model.Todo) error {
	result, err := r.DB.Exec(
		"UPDATE todos SET title=$1, completed=$2 WHERE id=$3 AND username=$4",
		todo.Title, todo.Completed, todo.Id, todo.Username,
	)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return errors.New("no rows updated")
	}

	return err
}

func (r *TodoRepository) Delete(id string) error {
	result, err := r.DB.Exec("DELETE FROM todos WHERE id=$1", id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return errors.New("no rows deleted")
	}

	return nil
}
