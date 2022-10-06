package main

import (
	"miniDiscord/chat"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "index.html")
	})

	mainHub := chat.NewHub("main", router)
	mainHub.AddRoom("general")
	mainHub.AddRoom("jokes")
	mainHub.AddRoom("memes")

	router.Run(":5050")
}
