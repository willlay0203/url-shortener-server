package api

import (
	"log"
	"net/http"
)

func HttpErrorHandler(w http.ResponseWriter, errorMsg string, statusCode int, err error) {
	if err != nil {
		log.Printf("%v", errorMsg)
		http.Error(w, errorMsg, statusCode)
	}
}
