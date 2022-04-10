package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func main() {
	server := gin.Default()
	server.LoadHTMLGlob("templates/*")
	server.Static("/assets", "./asset")
	server.GET("/login", LoginPage)
	server.Run(":8888")
}
