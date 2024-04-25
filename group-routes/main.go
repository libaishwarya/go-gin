package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	adminGroup := r.Group("/admin")

	adminGroup.GET("/users", func(c *gin.Context) {
		c.String(http.StatusOK, "Page to admin user")
	})

	adminGroup.GET("/roles", func(c *gin.Context) {
		c.String(http.StatusOK, "Page to admin roles")
	})

	adminGroup.GET("/policies", func(c *gin.Context) {
		c.String(http.StatusOK, "Page to admin policies")
	})

	log.Fatal(r.Run(":3000"))
}
