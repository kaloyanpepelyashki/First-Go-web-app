package controllers

import (
	"fmt"
	"net/http"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/kaloyanpepelyashki/task-manager-web-app/taskmanager"
)

type TasksController struct{}

func (tc *TasksController) CreateNewTaskHandler(ctx *gin.Context) {
	title := ctx.PostForm("title")
	description := ctx.PostForm("description")
	endDate := ctx.PostForm("endDate")

	layout := "2006-01-02"
	endDateParsed, err := time.Parse(layout, endDate)

	if err != nil {
		fmt.Println("error:", err)
		return
	}
	task := taskmanager.Task{}
	task.CreateTask(1, title, description, endDateParsed)

	// Redirect to a success page or any other desired action
	ctx.Redirect(http.StatusFound, "/task")

	fmt.Println("Task created successfully")

}
