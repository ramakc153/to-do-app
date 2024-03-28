package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"to-do-app/config"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type User struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var user User
	var userValid User
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	username, password, address, port, db_name := os.Getenv("USNAME"), os.Getenv("PASSWD"), os.Getenv("DB_IP"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME")
	db, err := config.Connect(username, password, address, port, db_name)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	err = c.ShouldBindJSON(&user)
	defer db.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("this is email ", user.Email)
	err = db.QueryRow("select * from users where email = ?", user.Email).Scan(&userValid.Email, &userValid.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Email not found",
		})
		return
	}

	if user.Password != userValid.Password {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Wrong Password",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Authorized",
	})
}
func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "this is from logoout endpoint",
	})
}
