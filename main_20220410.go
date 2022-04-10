package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HtmlData struct {
	Title string
	Body  string
}

func test(c *gin.Context) {
	data := new(HtmlData)
	data.Title = "首頁"
	data.Body = "GIN頁面"
	c.HTML(http.StatusOK, "index.html", data)
}

func main() {
	server := gin.Default()
	server.LoadHTMLGlob("templates/*")
	server.GET("/", test)
	server.Run(":8888")
}
