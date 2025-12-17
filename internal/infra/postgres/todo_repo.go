package postgres

import (
	"context"
	"todoDB/internal/domain/todo"

	"gorm.io/gorm"
)

// type TodoRepository interface {
// 	UpdateTodo(ctx context.Context, todoList *Task) error
// 	CreateTodo(ctx context.Context, todoList *Todo) error // empty
// 	DeleteTodoByID(ctx context.Context, todoListID int) error
// 	GetByID(ctx context.Context, todoLis *Todo) (*Todo, error)
// 	GetAllByUserName(ctx context.Context, userName string) (*Todo, error)
// }

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) todo.TodoRepository {
	return &todoRepository{db: db.Model(&todo.Todo{})}
}

func (t *todoRepository) UpdateTodo(ctx context.Context, todo *todo.Task) error {
	return t.db.WithContext(ctx).Save(todo).Error
}

func (t *todoRepository) CreateTodo(ctx context.Context, todo *todo.Todo) error {
	return t.db.WithContext(ctx).Create(todo).Error
}
