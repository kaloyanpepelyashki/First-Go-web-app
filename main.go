package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaloyanpepelyashki/task-manager-web-app/controllers"
	"github.com/kaloyanpepelyashki/task-manager-web-app/taskmanager"
)

func main() {
	server := gin.Default()
	server.Static("/static", "./static")
	server.LoadHTMLGlob("templates/*")
	TaskController := &controllers.TasksController{}

	server.GET("/task", func(ctx *gin.Context) {

		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Task Manager",
		})
	})
	server.GET("/task/create", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "create.html", gin.H{
			"title": "Create task",
		})
	})
	server.POST("/task/create", TaskController.CreateNewTaskHandler)

	server.GET("/task/taskslist", func(ctx *gin.Context) {
		fmt.Println("tasks list ==> ", taskmanager.TasksList)
		ctx.HTML(http.StatusOK, "taskslist.html", gin.H{
			"title":     "Tasks List",
			"tasksList": taskmanager.TasksList,
		})
	})

	server.GET("/task/completedtasks", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "completedtasks.html", gin.H{
			"title": "Completed tasks",
		})
	})

	server.Run(":8080")
}
