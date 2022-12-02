package main

import (
	"golang-mygram/app"

	"github.com/gin-gonic/gin"
)

func main() {
	app.StartDB()
	r := gin.Default()

	r.Run()
}
