package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// Todo list
type Todo struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAs time.Time `json:"createdAs"`
}

func main() {
	log.Println("Starting...")
	r := gin.Default()
	r.GET("/api/tasks", func(c *gin.Context) {
		todoList := []Todo{
			{ID: "1", Title: "Learn golang!", Completed: true, CreatedAs: time.Now()},
			{ID: "2", Title: "Learn docker!", Completed: true, CreatedAs: time.Now()},
			{ID: "3", Title: "Learn kafka!", Completed: false, CreatedAs: time.Now()},
		}
		c.JSON(200, todoList)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
