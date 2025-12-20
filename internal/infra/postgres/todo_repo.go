package postgres

import (
	"context"

	"todoDB/internal/domain/todo"

	"gorm.io/gorm"
)

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) todo.TodoRepository {
	return &todoRepository{db: db}
}

func (t *todoRepository) UpdateTodo(ctx context.Context, todo *todo.Todo) error {
	return t.db.WithContext(ctx).Save(todo).Error
}

func (t *todoRepository) CreateTodo(ctx context.Context, todo *todo.Todo) error {
	return t.db.WithContext(ctx).Create(todo).Error
}

func (t *todoRepository) DeleteTodoByID(ctx context.Context, todoListID int) error {
	return t.db.WithContext(ctx).Delete(&todo.Todo{}, todoListID).Error
}

func (t *todoRepository) GetByID(ctx context.Context, todoListID int) (*todo.Todo, error) {
	var theTodo todo.Todo
	err := t.db.WithContext(ctx).Preload("Tasks").First(&theTodo, todoListID).Error
	if err != nil {
		return nil, err
	} else {
		return &theTodo, nil
	}
}

func (t *todoRepository) GetAllByUserName(ctx context.Context, userName string) ([]*todo.Todo, error) {
	var todos []*todo.Todo
	err := t.db.WithContext(ctx).Preload("Tasks").Where("user_name = ?", userName).Find(&todos).Error
	if err != nil {
		return nil, err
	} else {
		return todos, nil
	}
}

func (t *todoRepository) AddTask(ctx context.Context, task *todo.Task) error {
	return t.db.WithContext(ctx).Create(task).Error
}

func (t *todoRepository) UpdateTask(ctx context.Context, task *todo.Task) error {
	return t.db.WithContext(ctx).Save(task).Error
}

func (t *todoRepository) DeleteTask(ctx context.Context, taskID int) error {
	return t.db.WithContext(ctx).Delete(&todo.Task{}, taskID).Error
}
