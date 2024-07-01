package api

import (
	"DMG/database"
	"DMG/structs"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Handlers for Client
func getClients(w http.ResponseWriter, r *http.Request) {
	var clients []structs.Client
	database.DB.Preload("Rides.Waypoints").Find(&clients)
	err := json.NewEncoder(w).Encode(clients)
	if err != nil {
		log.Fatal(err)
	}
}

func getClient(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var client structs.Client
	database.DB.Preload("Rides.Waypoints").First(&client, params["id"])
	if client.ID == 0 {
		http.Error(w, "Client not found", http.StatusNotFound)
		return
	}
	err := json.NewEncoder(w).Encode(client)
	if err != nil {
		log.Fatal(err)
	}
}

func createClient(w http.ResponseWriter, r *http.Request) {
	var client structs.Client
	_ = json.NewDecoder(r.Body).Decode(&client)
	database.DB.Create(&client)
	err := json.NewEncoder(w).Encode(client)
	if err != nil {
		log.Fatal(err)
	}
}

func updateClient(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var client structs.Client
	database.DB.First(&client, params["id"])
	if client.ID == 0 {
		http.Error(w, "Client not found", http.StatusNotFound)
		return
	}
	_ = json.NewDecoder(r.Body).Decode(&client)
	database.DB.Save(&client)
	err := json.NewEncoder(w).Encode(client)
	if err != nil {
		log.Fatal(err)
	}
}

func deleteClient(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var client structs.Client
	database.DB.First(&client, params["id"])
	if client.ID == 0 {
		http.Error(w, "Client not found", http.StatusNotFound)
		return
	}
	database.DB.Delete(&client)
	w.WriteHeader(http.StatusNoContent)
}

// Similar handlers can be implemented for Ride, Driver, and Waypoint
// Handlers for Ride
func getRides(w http.ResponseWriter, r *http.Request) {
	var rides []structs.Ride
	database.DB.Preload("Waypoints").Find(&rides)
	err := json.NewEncoder(w).Encode(rides)
	if err != nil {
		log.Fatal(err)
	}
}

func getRide(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var ride structs.Ride
	database.DB.Preload("Waypoints").First(&ride, params["id"])
	if ride.ID == 0 {
		http.Error(w, "Ride not found", http.StatusNotFound)
		return
	}
	err := json.NewEncoder(w).Encode(ride)
	if err != nil {
		log.Fatal(err)
	}
}

func createRide(w http.ResponseWriter, r *http.Request) {
	var ride structs.Ride
	_ = json.NewDecoder(r.Body).Decode(&ride)
	database.DB.Create(&ride)
	err := json.NewEncoder(w).Encode(ride)
	if err != nil {
		log.Fatal(err)
	}
}

func updateRide(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var ride structs.Ride
	database.DB.First(&ride, params["id"])
	if ride.ID == 0 {
		http.Error(w, "Ride not found", http.StatusNotFound)
		return
	}
	_ = json.NewDecoder(r.Body).Decode(&ride)
	database.DB.Save(&ride)
	err := json.NewEncoder(w).Encode(ride)
	if err != nil {
		log.Fatal(err)
	}
}

func deleteRide(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var ride structs.Ride
	database.DB.First(&ride, params["id"])
	if ride.ID == 0 {
		http.Error(w, "Ride not found", http.StatusNotFound)
		return
	}
	database.DB.Delete(&ride)
	w.WriteHeader(http.StatusNoContent)
}

// Handlers for Driver
func getDrivers(w http.ResponseWriter, r *http.Request) {
	var drivers []structs.Driver
	database.DB.Find(&drivers)
	err := json.NewEncoder(w).Encode(drivers)
	if err != nil {
		log.Fatal(err)
	}
}

func getDriver(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var driver structs.Driver
	database.DB.First(&driver, params["id"])
	if driver.ID == 0 {
		http.Error(w, "Driver not found", http.StatusNotFound)
		return
	}
	err := json.NewEncoder(w).Encode(driver)
	if err != nil {
		log.Fatal(err)
	}
}

func createDriver(w http.ResponseWriter, r *http.Request) {
	var driver structs.Driver
	_ = json.NewDecoder(r.Body).Decode(&driver)
	database.DB.Create(&driver)
	err := json.NewEncoder(w).Encode(driver)
	if err != nil {
		log.Fatal(err)
	}
}

func updateDriver(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var driver structs.Driver
	database.DB.First(&driver, params["id"])
	if driver.ID == 0 {
		http.Error(w, "Driver not found", http.StatusNotFound)
		return
	}
	_ = json.NewDecoder(r.Body).Decode(&driver)
	database.DB.Save(&driver)
	err := json.NewEncoder(w).Encode(driver)
	if err != nil {
		log.Fatal(err)
	}
}

func deleteDriver(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var driver structs.Driver
	database.DB.First(&driver, params["id"])
	if driver.ID == 0 {
		http.Error(w, "Driver not found", http.StatusNotFound)
		return
	}
	database.DB.Delete(&driver)
	w.WriteHeader(http.StatusNoContent)
}

// Handlers for Waypoint
func getWaypoints(w http.ResponseWriter, r *http.Request) {
	var waypoints []structs.Waypoint
	database.DB.Find(&waypoints)
	err := json.NewEncoder(w).Encode(waypoints)
	if err != nil {
		log.Fatal(err)
	}
}

func getWaypoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var waypoint structs.Waypoint
	database.DB.First(&waypoint, params["id"])
	if waypoint.ID == 0 {
		http.Error(w, "Waypoint not found", http.StatusNotFound)
		return
	}
	err := json.NewEncoder(w).Encode(waypoint)
	if err != nil {
		log.Fatal(err)
	}
}

func createWaypoint(w http.ResponseWriter, r *http.Request) {
	var waypoint structs.Waypoint
	_ = json.NewDecoder(r.Body).Decode(&waypoint)
	database.DB.Create(&waypoint)
	err := json.NewEncoder(w).Encode(waypoint)
	if err != nil {
		log.Fatal(err)
	}
}

func updateWaypoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var waypoint structs.Waypoint
	database.DB.First(&waypoint, params["id"])
	if waypoint.ID == 0 {
		http.Error(w, "Waypoint not found", http.StatusNotFound)
		return
	}
	_ = json.NewDecoder(r.Body).Decode(&waypoint)
	database.DB.Save(&waypoint)
	err := json.NewEncoder(w).Encode(waypoint)
	if err != nil {
		log.Fatal(err)
	}
}

func deleteWaypoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var waypoint structs.Waypoint
	database.DB.First(&waypoint, params["id"])
	if waypoint.ID == 0 {
		http.Error(w, "Waypoint not found", http.StatusNotFound)
		return
	}
	database.DB.Delete(&waypoint)
	w.WriteHeader(http.StatusNoContent)
}
