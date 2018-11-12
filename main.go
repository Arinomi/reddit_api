package main

import (
	"log"
	"net/http"
	"os"
	"reddit_api/api"

	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Making oauth for the api and setting up a session
	api.InitAuth()

	// Set up handlers
	r := mux.NewRouter()

	// first handlers
	r.HandleFunc("/reddit", api.Redirect).Methods("GET")
	r.HandleFunc("/reddit/api/", api.InfoHandler).Methods("GET")
	r.HandleFunc("/reddit/api/submission/", api.SubmissionHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":"+port, r))
}