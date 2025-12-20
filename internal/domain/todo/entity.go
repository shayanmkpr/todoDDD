package todo

type Todo struct {
	TodoID   int `json:"todo_id" binding:"required"`
	Tasks    []*Task
	Title    string
	UserName string `json:"user_name" binding:"required"`
}

type Task struct {
	TaksID int    `json:"task_id" binding:"required"`
	Title  string `json:"title"`
	Status string `json:"status"`
	Body   string `json:"body"`
}
