package main

import (
	"log"
	"mirzaadr/calculator-api/cmd/api"
)

func main() {
	server := api.NewAPIServer(":4000")
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
