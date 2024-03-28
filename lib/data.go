package lib

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
	Body      string `json:"body"`
}

func EncodedData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	peter := []User{
		{
			Firstname: "John",
			Lastname:  "Doe",
			Age:       25,
			Body:      "GGS",
		},
		{
			Firstname: "John",
			Lastname:  "Doe",
			Age:       25,
			Body:      "ayayay",
		},
	}
	json.NewEncoder(w).Encode(peter)
}

func GetUsers(c *gin.Context) {
	peter := []User{
		{
			Firstname: "John",
			Lastname:  "Doe",
			Age:       25,
			Body:      "GGS",
		},
		{
			Firstname: "John",
			Lastname:  "Doe",
			Age:       25,
			Body:      "ayayay",
		},
	}
	c.JSON(200, gin.H{
		"body": peter,
	})
}
