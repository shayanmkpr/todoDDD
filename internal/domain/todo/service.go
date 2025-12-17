package todo

import "errors"

func (t *Todo) AddTask(newTask *Task) error {
	if newTask.Body == "" {
		return errors.New("the new task is empty")
	}
	t.Tasks = append(t.Tasks, newTask)
	return nil
}

func (t *Todo) RemoveTaskByID(taskID int) error {
	for i, task := range t.Tasks {
		if task.TaksID == taskID {
			t.Tasks = append(t.Tasks[:i], t.Tasks[:i+1]...)
			return nil
		}
	}
	return errors.New("no Such task was found in the todo list")
}
