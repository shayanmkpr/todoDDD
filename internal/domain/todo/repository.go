package todo

import "context"

type TodoRepository interface {
	UpdateTodo(ctx context.Context, todo *Todo) error
	CreateTodo(ctx context.Context, todo *Todo) error
	DeleteTodoByID(ctx context.Context, todoListID int) error
	GetByID(ctx context.Context, todoListID int) (*Todo, error)
	GetAllByUserName(ctx context.Context, userName string) ([]*Todo, error)
	AddTask(ctx context.Context, task *Task) error
	UpdateTask(ctx context.Context, task *Task) error
	DeleteTask(ctx context.Context, taskID int) error
}

//
// type TaskRepo interface {
// 	UpdateTask(ctx context.Context, t *Task) error
// 	CreateTask(ctx context.Context, t *Task) error
// 	DeleteTask(ctx context.Context, t *Task) error
// 	GetByID(ctx context.Context, taskID int) (*Task, error)
// 	GetAllByTitle(ctx context.Context, taskTitle string) ([]*Task, error)
// 	GetAllByStatus(ctx context.Context, taskStatus string) ([]*Task, error)
// }
