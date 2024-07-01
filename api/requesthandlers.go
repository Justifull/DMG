package api

import (
	"DMG/database"
	"DMG/structs"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// Handlers for Client
func getClients(w http.ResponseWriter, r *http.Request) {
	var clients []structs.Client
	database.DB.Preload("Rides.Waypoints").Find(&clients)
	json.NewEncoder(w).Encode(clients)
}

func getClient(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var client structs.Client
	database.DB.Preload("Rides.Waypoints").First(&client, params["id"])
	if client.ID == 0 {
		http.Error(w, "Client not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(client)
}

func createClient(w http.ResponseWriter, r *http.Request) {
	var client structs.Client
	_ = json.NewDecoder(r.Body).Decode(&client)
	database.DB.Create(&client)
	json.NewEncoder(w).Encode(client)
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
	json.NewEncoder(w).Encode(client)
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
	json.NewEncoder(w).Encode(rides)
}

func getRide(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var ride structs.Ride
	database.DB.Preload("Waypoints").First(&ride, params["id"])
	if ride.ID == 0 {
		http.Error(w, "Ride not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(ride)
}

func createRide(w http.ResponseWriter, r *http.Request) {
	var ride structs.Ride
	_ = json.NewDecoder(r.Body).Decode(&ride)
	database.DB.Create(&ride)
	json.NewEncoder(w).Encode(ride)
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
	json.NewEncoder(w).Encode(ride)
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
	json.NewEncoder(w).Encode(drivers)
}

func getDriver(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var driver structs.Driver
	database.DB.First(&driver, params["id"])
	if driver.ID == 0 {
		http.Error(w, "Driver not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(driver)
}

func createDriver(w http.ResponseWriter, r *http.Request) {
	var driver structs.Driver
	_ = json.NewDecoder(r.Body).Decode(&driver)
	database.DB.Create(&driver)
	json.NewEncoder(w).Encode(driver)
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
	json.NewEncoder(w).Encode(driver)
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
	json.NewEncoder(w).Encode(waypoints)
}

func getWaypoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var waypoint structs.Waypoint
	database.DB.First(&waypoint, params["id"])
	if waypoint.ID == 0 {
		http.Error(w, "Waypoint not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(waypoint)
}

func createWaypoint(w http.ResponseWriter, r *http.Request) {
	var waypoint structs.Waypoint
	_ = json.NewDecoder(r.Body).Decode(&waypoint)
	database.DB.Create(&waypoint)
	json.NewEncoder(w).Encode(waypoint)
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
	json.NewEncoder(w).Encode(waypoint)
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
