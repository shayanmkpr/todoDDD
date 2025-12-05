package todo

import "context"

type TodoRepo interface {
	UpdateTodo(ctx context.Context, t *Todo) error

	CreateTodo(ctx context.Context, t *Todo) error

	DeleteTodo(ctx context.Context, t *Todo) error

	GetAllTodo(ctx context.Context, t *Todo) ([]*Todo, error)

	GetByID(ctx context.Context, t *Todo) (*Todo, error)
}
