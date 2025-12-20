package todo

type Todo struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Title    string `json:"title"`
	UserName string `json:"user_name" gorm:"index"`
	Tasks    []*Task `json:"tasks" gorm:"foreignKey:TodoID"`
}

type Task struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	TodoID uint   `json:"todo_id"`
	Title  string `json:"title"`
	Status string `json:"status"`
	Body   string `json:"body"`
}
