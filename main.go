package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	// Create a new router
	r := mux.NewRouter()

	// Define your HTTP routes using the router
	r.HandleFunc("/user", api.createUserHandler).Methods("POST")
	r.HandleFunc("/user/{id}", api.getUserHandler).Methods("GET")
	r.HandleFunc("/user/{id}", api.updateUserHandler).Methods("PUT")
	r.HandleFunc("/user/{id}", api.deleteUserHandler).Methods("DELETE")

	// Start the HTTP server on port 8090
	log.Println("Server listening on :8090")
	log.Fatal(http.ListenAndServe(":8090", r))
}
