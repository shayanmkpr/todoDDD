package todo

import "errors"

func (t *Todo) AddTask(newTask *Task) error {
	if newTask.Body == "" {
		return errors.New("the new task is empty")
	}
	t.Tasks = append(t.Tasks, newTask)
	return nil
}

func (t *Todo) RemoveTaskByID(taskID uint) error {
	for i, task := range t.Tasks {
		if task.ID == taskID {
			// Remove element at index i
			t.Tasks = append(t.Tasks[:i], t.Tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("no such task was found in the todo list")
}
