package todo

type Todo struct {
	TodoID   int `json:"todo_id"`
	Tasks    []*Task
	userName int `json:"user_name"`
}

type Task struct {
	TaksID int    `json:"task_id" binding:"required"`
	Title  string `json:"title"`
	Status string `json:"status"`
	Body   string `json:"body"`
}
