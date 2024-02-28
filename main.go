package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const PORT = "8080"

func main() {
	router := httprouter.New()

	fmt.Printf("Server listening on %v", PORT)
	log.Fatal(http.ListenAndServe(":8080", router))
}
