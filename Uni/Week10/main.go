package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type toDo struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Status string  `json:"status"` // in process, ongoing, done
	Importance  string `json:"Importance"` // high, low, medium
}

var todo = []toDo{
	{ID: "1", Title: "Go to shop", Status: "ongoing", Importance: "medium"},
	{ID: "2", Title: "Do homework", Status: "in process", Importance: "high"},
	{ID: "3", Title: "Prepare to a seminar", Status: "in process", Importance: "high"},
	{ID: "4", Title: "Read scientific articles", Status: "ongoing", Importance: "medium"},
	{ID: "5", Title: "Book a manicure", Status: "ongoing", Importance: "low"},
	{ID: "6", Title: "Feed a cat", Status: "done", Importance: "high"},
}

func getWhatToDos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todo)
}

func postWhatToDo(c *gin.Context) {
	var what_todo toDo

	if err := c.BindJSON(&what_todo); err != nil {
		return
	}

	todo = append(todo, what_todo)
	c.IndentedJSON(http.StatusCreated, what_todo)
}


// func getToDoByID(c *gin.Context) {
// 	id := c.Param("id")

// 	for _, tDo := range todo {
// 		if tDo.ID == id {
// 			c.IndentedJSON(http.StatusOK, tDo)
// 			return
// 		}
// 	}

// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "That action not found! Maybe you've done this. If so, Good for you)"})
// }

func getToDoByImportance(c *gin.Context) {
	importance := c.Param("Importance")

	for _, tDo := range todo {
		if tDo.Importance == importance {
			c.IndentedJSON(http.StatusOK, tDo)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "That action not found! Maybe you've done this. If so, Good for you)"})
}


func deleteTodoByID(c *gin.Context) {
	id := c.Param("id")

	for i, tDO := range todo {
		if tDO.ID == id {
			todo = append(todo[:i], todo[i+1:]...)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "That action not found"})
}

func main() {
	router := gin.Default()
	router.GET("/todos", getWhatToDos)
	router.POST("/postTodo", postWhatToDo)
	// router.GET("/todos/:id", getToDoByID)
	router.GET("/todos/:Importance", getToDoByImportance)
	router.DELETE("/deleteTodo/:id", deleteTodoByID)

	router.Run("localhost:8080")
}
