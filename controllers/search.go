package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SearchByName(c *gin.Context) {
	name := c.Query("name")
	message := fmt.Sprintf("book name: %v", name)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": message,
	})
}
