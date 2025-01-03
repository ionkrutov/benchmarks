package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.GET("/hello", func(c *gin.Context) {
		c.Writer.Write([]byte("Hello world!\n"))
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
