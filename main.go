package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Task struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

var tasks []Task

func main() {
	router := gin.Default()

	// Загружаем шаблоны HTML
	router.SetHTMLTemplate(template.Must(template.ParseFiles("templates/index.html")))

	// Маршрут для главной страницы
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// Маршрут для получения списка задач
	router.GET("/tasks", func(c *gin.Context) {
		c.JSON(http.StatusOK, tasks)
	})

	// Маршрут для добавления новой задачи
	router.POST("/tasks", func(c *gin.Context) {
		var task Task
		if err := c.BindJSON(&task); err != nil {
			log.Println("Error adding task:", err)
			c.Status(http.StatusBadRequest)
			return
		}

		task.ID = len(tasks) + 1
		tasks = append(tasks, task)

		c.JSON(http.StatusOK, task)
	})

	// Маршрут для изменения статуса задачи
	router.PATCH("/tasks/:id/status", func(c *gin.Context) {
		taskID := c.Param("id")
		newStatus := c.Query("status")

		for i, task := range tasks {
			if strconv.Itoa(task.ID) == taskID {
				tasks[i].Status = newStatus
				break
			}
		}

		c.Status(http.StatusOK)
	})

	// Запуск веб-сервера
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
