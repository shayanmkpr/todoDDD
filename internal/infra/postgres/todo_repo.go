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
	return &todoRepository{db: db.Model(&todo.Todo{})}
}

func (t *todoRepository) UpdateTodo(ctx context.Context, todo *todo.Task) error { // should check this update and save logic relation ship
	return t.db.WithContext(ctx).Save(todo).Error
}

func (t *todoRepository) CreateTodo(ctx context.Context, todo *todo.Todo) error { // just makes an empty todo. not sure if I want to insert some tasks as well. maybe a single one.
	return t.db.WithContext(ctx).Create(todo).Error
}

func (t *todoRepository) DeleteTodoByID(ctx context.Context, todoListID int) error {
	return t.db.WithContext(ctx).Where("todo_id = ?", todoListID).Delete(&todo.Todo{}).Error
}

func (t *todoRepository) GetByID(ctx context.Context, todoList *todo.Todo) (*todo.Todo, error) {
	var theTodo todo.Todo
	err := t.db.WithContext(ctx).Where("id = ?", todoList.TodoID).First(&theTodo).Error
	if err != nil {
		return nil, err
	} else {
		return &theTodo, nil
	}
}

func (t *todoRepository) GetAllByUserName(ctx context.Context, userName string) (*todo.Todo, error) {
	var theTodo todo.Todo
	err := t.db.WithContext(ctx).Where("user_name = ?", userName).First(&theTodo).Error
	if err != nil {
		return nil, err
	} else {
		return &theTodo, nil
	}
}
