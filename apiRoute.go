package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justiruel/penggajianPaperbag/api"
)

// func RegisterRoute(r *mux.Router) {
func RegisterRoute() {
	// Create a new router
	r := mux.NewRouter()

	// Define your HTTP routes using the router
	r.HandleFunc("/user", api.CreateUserHandler).Methods("POST")
	r.HandleFunc("/user/{id}", api.GetUserHandler).Methods("GET")
	r.HandleFunc("/user/{id}", api.UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/user/{id}", api.DeleteUserHandler).Methods("DELETE")

	// Start the HTTP server on port 8090
	log.Println("Server listening on :8090")
	log.Fatal(http.ListenAndServe(":8090", r))
}
