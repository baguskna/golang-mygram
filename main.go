package main

import (
	"golang-mygram/app"
	"golang-mygram/router"
)

func main() {
	app.StartDB()
	r := router.StartApp()
	r.Run()
}
