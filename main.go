package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// var f embed.FS
	r := gin.Default()

	// r.StaticFile("/", "./public/index.html")
	// r.Static("/public", "./public")
	// r.StaticFS("/fs", http.FileSystem(http.FS(f)))

	r.GET("/employee", func(c *gin.Context) {
		c.File("./public/employee.html")
	})

	r.POST("/employee", func(c *gin.Context) {
		c.String(http.StatusOK, "New request posted successfully")
	})

	log.Fatal(r.Run(":3000"))
}
