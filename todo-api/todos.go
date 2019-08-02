package main

import (
	"math/rand"
	"strconv"
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

// TodoRouter for all todo
func TodoRouter(router *gin.Engine) {
	todoList := []Todo{
		{ID: "1", Title: "Learn golang!", Completed: true, CreatedAs: time.Now()},
		{ID: "2", Title: "Learn docker!", Completed: true, CreatedAs: time.Now()},
		{ID: "3", Title: "Learn kafka!", Completed: false, CreatedAs: time.Now()},
	}
	router.GET("/api/tasks", func(c *gin.Context) {
		todo := Todo{ID: strconv.Itoa(rand.Int()), Title: "Learn SpringBoot!", Completed: true, CreatedAs: time.Now()}
		todoList = append(todoList, todo)
		c.JSON(200, todoList)
	})
}
