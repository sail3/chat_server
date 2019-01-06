package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "status ok!!",
		})
	})
	r.GET("/ws", webHandler)
	r.Run(":8080")
}

func webHandler(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"message": "hola mundo",
	})
	con, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(con)
}
