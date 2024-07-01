package api

import (
	"DMG/database"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var (
	ROUTER *mux.Router // Storing the current router for better access
)

// StartAPI initializes the api and its handlers
func StartAPI() {
	// Init the database
	err := database.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}

	// Init a new router acting as a request multiplexer. It handles all registered
	// routes from the HTTP requests and executes the corresponding handlers
	ROUTER = mux.NewRouter()

	// Client routes
	ROUTER.HandleFunc("/clients", getClients).Methods("GET")
	ROUTER.HandleFunc("/clients/{id}", getClient).Methods("GET")
	ROUTER.HandleFunc("/clients", createClient).Methods("POST")
	ROUTER.HandleFunc("/clients/{id}", updateClient).Methods("PUT")
	ROUTER.HandleFunc("/clients/{id}", patchClient).Methods("PATCH")
	ROUTER.HandleFunc("/clients/{id}", deleteClient).Methods("DELETE")

	// Ride routes
	ROUTER.HandleFunc("/rides", getRides).Methods("GET")
	ROUTER.HandleFunc("/rides/{id}", getRide).Methods("GET")
	ROUTER.HandleFunc("/rides", createRide).Methods("POST")
	ROUTER.HandleFunc("/rides/{id}", updateRide).Methods("PUT")
	ROUTER.HandleFunc("/rides/{id}", patchRide).Methods("PATCH")
	ROUTER.HandleFunc("/rides/{id}", deleteRide).Methods("DELETE")

	// Driver routes
	ROUTER.HandleFunc("/drivers", getDrivers).Methods("GET")
	ROUTER.HandleFunc("/drivers/{id}", getDriver).Methods("GET")
	ROUTER.HandleFunc("/drivers", createDriver).Methods("POST")
	ROUTER.HandleFunc("/drivers/{id}", updateDriver).Methods("PUT")
	ROUTER.HandleFunc("/drivers/{id}", patchDriver).Methods("PATCH")
	ROUTER.HandleFunc("/drivers/{id}", deleteDriver).Methods("DELETE")

	// Waypoint routes
	ROUTER.HandleFunc("/waypoints", getWaypoints).Methods("GET")
	ROUTER.HandleFunc("/waypoints/{id}", getWaypoint).Methods("GET")
	ROUTER.HandleFunc("/waypoints", createWaypoint).Methods("POST")
	ROUTER.HandleFunc("/waypoints/{id}", updateWaypoint).Methods("PUT")
	ROUTER.HandleFunc("/waypoints/{id}", patchWaypoint).Methods("PATCH")
	ROUTER.HandleFunc("/waypoints/{id}", deleteWaypoint).Methods("DELETE")

	fmt.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", ROUTER))
}
