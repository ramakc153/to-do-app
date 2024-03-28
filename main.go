package main

import (
	"fmt"
	"net/http"
	"os"
	"to-do-app/controllers"
	"to-do-app/lib"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func index(c *gin.Context) {
	c.Data(http.StatusOK, "text/plain; charset=utf-8", []byte("dlkldaskdsa"))

}

func hello(c *gin.Context) {
	c.Data(http.StatusOK, "text/html", []byte("Hello worlds"))
}
func data(c *gin.Context) {
	data := map[string]string{
		"Name": "John Wick",
		"Age":  "31",
	}
	c.HTML(http.StatusOK, "template.html", data)
}

func main() {
	godotenv.Load()

	ip := os.Getenv("IP")
	addr := fmt.Sprintf("%v:5000", ip)

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/path", lib.GetUsers)
	r.GET("/index", index)
	r.GET("/hello", hello)
	r.GET("/login", controllers.Login)
	r.GET("/logout", controllers.Logout)
	r.GET("/data", data)
	r.GET("/note/:id", controllers.GetNoteById)
	r.GET("/note", controllers.SearchByName)
	r.POST("/note", controllers.PostNote)
	r.Run(addr)

}
