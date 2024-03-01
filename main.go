package main

import (
	"log"
	"server/api"
)

func main() {
	const PORT string = ":8080"

	server := api.Server{Port: PORT}

	log.Fatal(server.Start())
}
