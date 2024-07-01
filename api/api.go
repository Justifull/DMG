package api

import (
	"DMG/database"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func StartAPI() {
	err := database.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()

	// Client routes
	router.HandleFunc("/clients", getClients).Methods("GET")
	router.HandleFunc("/clients/{id}", getClient).Methods("GET")
	router.HandleFunc("/clients", createClient).Methods("POST")
	router.HandleFunc("/clients/{id}", updateClient).Methods("PUT")
	router.HandleFunc("/clients/{id}", deleteClient).Methods("DELETE")

	// Ride routes
	router.HandleFunc("/rides", getRides).Methods("GET")
	router.HandleFunc("/rides/{id}", getRide).Methods("GET")
	router.HandleFunc("/rides", createRide).Methods("POST")
	router.HandleFunc("/rides/{id}", updateRide).Methods("PUT")
	router.HandleFunc("/rides/{id}", deleteRide).Methods("DELETE")

	// Driver routes
	router.HandleFunc("/drivers", getDrivers).Methods("GET")
	router.HandleFunc("/drivers/{id}", getDriver).Methods("GET")
	router.HandleFunc("/drivers", createDriver).Methods("POST")
	router.HandleFunc("/drivers/{id}", updateDriver).Methods("PUT")
	router.HandleFunc("/drivers/{id}", deleteDriver).Methods("DELETE")

	// Waypoint routes
	router.HandleFunc("/waypoints", getWaypoints).Methods("GET")
	router.HandleFunc("/waypoints/{id}", getWaypoint).Methods("GET")
	router.HandleFunc("/waypoints", createWaypoint).Methods("POST")
	router.HandleFunc("/waypoints/{id}", updateWaypoint).Methods("PUT")
	router.HandleFunc("/waypoints/{id}", deleteWaypoint).Methods("DELETE")

	fmt.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
