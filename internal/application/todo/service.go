package application

import (
	"context"
	"todoDB/internal/domain/auth"
	"todoDB/internal/domain/todo"
)

//remove task from todo, get those by status.

type TodoServices struct {
	todoRepo todo.TodoRepository
	authRepo auth.AuthenticationRepo // since it should validate the tokens and authienciate stuff
}

func NewTodoService(t todo.TodoRepository, a auth.AuthenticationRepo) *TodoServices {
	return &TodoServices{
		todoRepo: t,
		authRepo: a,
	}
}

func (s *TodoServices) CreateTodo(ctx context.Context, userName string, firstTask *todo.Task) error {
	err := s.todoRepo.CreateTodo(ctx, &todo.Todo{UserName: userName})
	if err != nil {
		return err
	}
	return nil
}

func (s *TodoServices) RemoveTodo(ctx context.Context, todoID int) error {
	err := s.todoRepo.DeleteTodoByID(ctx, todoID)
	if err != nil {
		return err
	}
	return nil
}

func (s *TodoServices) GetTodoByUserName(ctx context.Context, userName string) (*todo.Todo, error) {
	theTodo, err := s.todoRepo.GetAllByUserName(ctx, userName)
	if err != nil {
		return &todo.Todo{}, err
	}
	return theTodo, nil
}

func (s *TodoServices) AddTask(ctx context.Context, todoID int, task *todo.Task) error {
	return nil
}

func (s *TodoServices) UpdateTask(ctx context.Context, task *todo.Task) error {
	return nil
}

func (s *TodoServices) DeleteTask(ctx context.Context, taskID int) error {
	return nil
}
