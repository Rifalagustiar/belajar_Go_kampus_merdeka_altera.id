package main

import (
	"crud/config"
	"crud/route"
)

// get all users

func main() {
	// create a new echo instance
	config.Init()
	e := route.New()
	e.Logger.Fatal(e.Start(":8080"))
}
