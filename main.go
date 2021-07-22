package main

import (
	"os"

	"github.com/go-sol-sum/router"
)

//Instruction os how to run the project on README
func main() {
	router := router.SetupRouter()

	router.Run(os.Getenv("PORT"))
}
