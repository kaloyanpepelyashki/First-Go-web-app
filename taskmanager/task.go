package taskmanager

import (
	"fmt"
	"time"
)

type Task struct {
	Id          int
	Title       string
	Description string
	CreatedAt   time.Time
	EndDate     time.Time
	Completed   bool
	Deleted     bool
	DeleteDate  *time.Time
}

var TasksList []Task

func (t *Task) CreateTask(id int, title string, description string, endDate time.Time) {
	task := Task{
		Id:          id,
		Title:       title,
		Description: description,
		CreatedAt:   time.Now(),
		EndDate:     endDate,
		Completed:   false,
		Deleted:     false,
	}

	TasksList = append(TasksList, task)
	fmt.Println(TasksList)
}
