package main

import (
	c "lgin/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	rExample := r.Group("/example")
	rExample.GET("/:id", c.GetExample)

	// rUsers := r.Group("/users")

	r.Run()
}
