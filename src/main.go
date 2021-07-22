package main

import "github.com/go-sol-sum/router"

func main() {
	router := router.SetupRouter()

	router.Run()
}
