package service

import (
	"errors"
	"todo-api/internal/model"
	"todo-api/internal/repository"

	"github.com/google/uuid"
)

type TodoService struct {
	TodoRepo *repository.TodoRepository
}

func NewTodoService(todoRepo *repository.TodoRepository) *TodoService {
	return &TodoService{
		TodoRepo: todoRepo,
	}
}

func (s *TodoService) Create(todo model.Todo) error {
	if todo.Title == "" {
		return errors.New("title is required")
	}
	if todo.Username == "" {
		return errors.New("username is required")
	}
	todo.Id = uuid.NewString()
	return s.TodoRepo.Create(todo)
}

func (s *TodoService) FindByUsername(username string) ([]model.Todo, error) {
	return s.TodoRepo.FindByUsername(username)
}

func (s *TodoService) Update(todo model.Todo) error {
	return s.TodoRepo.Update(todo)
}

func (s *TodoService) Delete(id string) error {
	return s.TodoRepo.Delete(id)
}
