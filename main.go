package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/employees/:username/*rest", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"username": c.Param("username"),
			"rest":     c.Param("rest"),
		})
	})

	log.Fatal(r.Run(":3000"))
}
