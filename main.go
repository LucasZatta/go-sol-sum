package main

import (
	"os"

	"github.com/go-sol-sum/router"
)

func main() {
	router := router.SetupRouter()

	router.Run(os.Getenv("PORT"))
}
