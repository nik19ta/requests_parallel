package main

import (
	"log"
	"req_parallel/server"
)

func main() {

	app := server.NewApp()

	if err := app.Run("3069"); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
