package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"server/lib"
	"server/types"

	"github.com/rs/cors"
)

type Server struct {
	Port string
}

func (s *Server) Start() error {
	router := http.NewServeMux()

	// Defining our routes
	router.HandleFunc("GET /healthcheck", healthCheck)
	// Endpoint to shorten url
	router.HandleFunc("POST /url", shortenUrl)
	// Endpoint to hit shorten url
	router.HandleFunc("GET /url/{shortLink}", redirectLink)
	fmt.Printf("Server listening on %v\n", s.Port)

	// CORS middleware
	handler := cors.Default().Handler(router)
	return http.ListenAndServe(s.Port, handler)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Server is OK")
}

func shortenUrl(w http.ResponseWriter, r *http.Request) {
	var err error
	// Reads body into a byte array
	body, err := io.ReadAll(r.Body)

	// Checks for error
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
	}

	var url types.Url

	// Treat this like an object mapper
	err = json.Unmarshal(body, &url)

	if err != nil {
		log.Printf("Error mapping request body: %v", err)
		http.Error(w, "Error mapping request body", http.StatusBadRequest)
	}

	// Log and return input
	log.Printf("%v", string(body))

	// Process the url passed in
	// Randomise
	code := lib.RandomString()

	// Send request to POST to db
	err = postNewShortenedUrl(code, url.Uri)

	if err != nil {
		log.Printf("Error posting to DB")
		http.Error(w, "Error mapping request body", http.StatusInternalServerError)
	}

	// DB has been posted return value to user
	type Payload struct {
		NewLink string
	}

	// Send back json payload
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&Payload{r.Host + "/url/" + code})
}

func redirectLink(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("shortLink")

	link, err := getUrl(code)

	log.Println(link)

	if err != nil {
		http.Error(w, "Link not matched", http.StatusNotFound)
	}

	http.Redirect(w, r, link, http.StatusMultipleChoices)
}
