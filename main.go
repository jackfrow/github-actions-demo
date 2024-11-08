package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {

	fmt.Println("test04")

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "hello actions")
	})
	r.GET("/hello", func(c *gin.Context) {
		log.Print("hello request")

		c.JSON(200, gin.H{
			"code":  200,
			"state": "success",
			"data":  "hello word",
		})
	})

	r.Run(":9090")

}
