package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetNoteById(c *gin.Context) {
	id := c.Param("id")
	message := fmt.Sprintf("book with id: %v", id)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"id":      id,
		"message": message,
	})
}

type Note struct {
	Id    string `json:"id" binding:"required"`
	Title string `json:"title" binding:"required"`
}

var listOfNotes []Note

func PostNote(c *gin.Context) {
	var notePost Note
	err := c.ShouldBindJSON(&notePost)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  http.StatusUnprocessableEntity,
			"message": "Field cannot be blank",
		})
		return
	}

	listOfNotes = append(listOfNotes, notePost)

	c.JSON(http.StatusOK, listOfNotes)
}
