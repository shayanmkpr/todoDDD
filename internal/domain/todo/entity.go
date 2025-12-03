package todo

type Todo struct {
	Tasks []*Tasks
}

type Tasks struct {
	Title  string
	Status string
	body   string
}
