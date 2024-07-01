package api

import (
	"DMG/database"
	"DMG/structs"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

/*********************************************************************/
/************************** REQUEST HANDLERS *************************/
/*********************************************************************/

// Clients get-request handler
func getClients(w http.ResponseWriter, r *http.Request) {
	var clients []structs.Client
	if err := database.DB.Preload("Rides.Waypoints").Find(&clients).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(clients); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Client get-request handler
func getClient(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var client structs.Client
	if err := database.DB.Preload("Rides.Waypoints").First(&client, params["id"]).Error; err != nil {
		http.Error(w, "Error 404: Client not found", http.StatusNotFound)
		return
	}
	if err := json.NewEncoder(w).Encode(client); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Client post-request handler
func createClient(w http.ResponseWriter, r *http.Request) {
	var client structs.Client
	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := database.DB.Create(&client).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(client); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Client put-request handler
func updateClient(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var client structs.Client
	if err := database.DB.First(&client, params["id"]).Error; err != nil {
		http.Error(w, "Error 404: Client not found", http.StatusNotFound)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := database.DB.Save(&client).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(client); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Client patch-request handler
func patchClient(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var client structs.Client
	if err := database.DB.First(&client, params["id"]).Error; err != nil {
		http.Error(w, "Error 404: Client not found", http.StatusNotFound)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := database.DB.Model(&client).Updates(client).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(client); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Client delete-request handler
func deleteClient(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var client structs.Client
	if err := database.DB.First(&client, params["id"]).Error; err != nil {
		http.Error(w, "Error 404: Client not found", http.StatusNotFound)
		return
	}
	if err := database.DB.Delete(&client).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// Rides get-request handler
func getRides(w http.ResponseWriter, r *http.Request) {
	var rides []structs.Ride
	if err := database.DB.Preload("Waypoints").Find(&rides).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(rides); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Ride get-request handler
func getRide(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var ride structs.Ride
	if err := database.DB.Preload("Waypoints").First(&ride, params["id"]).Error; err != nil {
		http.Error(w, "Error 404: Ride not found", http.StatusNotFound)
		return
	}
	if err := json.NewEncoder(w).Encode(ride); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Ride post-request handler
func createRide(w http.ResponseWriter, r *http.Request) {
	var ride structs.Ride
	if err := json.NewDecoder(r.Body).Decode(&ride); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := database.DB.Create(&ride).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(ride); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Ride put-request handler
func updateRide(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var ride structs.Ride
	if err := database.DB.First(&ride, params["id"]).Error; err != nil {
		http.Error(w, "Error 404: Ride not found", http.StatusNotFound)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&ride); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := database.DB.Save(&ride).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(ride); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Ride patch-request handler
func patchRide(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var ride structs.Ride
	if err := database.DB.First(&ride, params["id"]).Error; err != nil {
		http.Error(w, "Error 404: Ride not found", http.StatusNotFound)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&ride); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := database.DB.Model(&ride).Updates(ride).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(ride); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Ride delete-request handler
func deleteRide(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var ride structs.Ride
	if err := database.DB.First(&ride, params["id"]).Error; err != nil {
		http.Error(w, "Error 404: Ride not found", http.StatusNotFound)
		return
	}
	if err := database.DB.Delete(&ride).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// Drivers get-request handler
func getDrivers(w http.ResponseWriter, r *http.Request) {
	var drivers []structs.Driver
	if err := database.DB.Find(&drivers).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(drivers); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Driver get-request handler
func getDriver(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var driver structs.Driver
	if err := database.DB.First(&driver, params["id"]).Error; err != nil {
		http.Error(w, "Error 404: Driver not found", http.StatusNotFound)
		return
	}
	if err := json.NewEncoder(w).Encode(driver); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Driver post-request handler
func createDriver(w http.ResponseWriter, r *http.Request) {
	var driver structs.Driver
	if err := json.NewDecoder(r.Body).Decode(&driver); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := database.DB.Create(&driver).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(driver); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Driver put-request handler
func updateDriver(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var driver structs.Driver
	if err := database.DB.First(&driver, params["id"]).Error; err != nil {
		http.Error(w, "Error 404: Driver not found", http.StatusNotFound)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&driver); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := database.DB.Save(&driver).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(driver); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Driver patch-request handler
func patchDriver(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var driver structs.Driver
	if err := database.DB.First(&driver, params["id"]).Error; err != nil {
		http.Error(w, "Error 404: Driver not found", http.StatusNotFound)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&driver); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := database.DB.Model(&driver).Updates(driver).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(driver); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Driver delete-request handler
func deleteDriver(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var driver structs.Driver
	if err := database.DB.First(&driver, params["id"]).Error; err != nil {
		http.Error(w, "Error 404: Driver not found", http.StatusNotFound)
		return
	}
	if err := database.DB.Delete(&driver).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// Waypoints get-request handler
func getWaypoints(w http.ResponseWriter, r *http.Request) {
	var waypoints []structs.Waypoint
	if err := database.DB.Find(&waypoints).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(waypoints); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Waypoint get-request handler
func getWaypoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var waypoint structs.Waypoint
	if err := database.DB.First(&waypoint, params["id"]).Error; err != nil {
		http.Error(w, "Error 404: Waypoint not found", http.StatusNotFound)
		return
	}
	if err := json.NewEncoder(w).Encode(waypoint); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Waypoint post-request handler
func createWaypoint(w http.ResponseWriter, r *http.Request) {
	var waypoint structs.Waypoint
	if err := json.NewDecoder(r.Body).Decode(&waypoint); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := database.DB.Create(&waypoint).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(waypoint); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Waypoint put-request handler
func updateWaypoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var waypoint structs.Waypoint
	if err := database.DB.First(&waypoint, params["id"]).Error; err != nil {
		http.Error(w, "Error 404: Waypoint not found", http.StatusNotFound)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&waypoint); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := database.DB.Save(&waypoint).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(waypoint); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Waypoint patch-request handler
func patchWaypoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var waypoint structs.Waypoint
	if err := database.DB.First(&waypoint, params["id"]).Error; err != nil {
		http.Error(w, "Error 404: Waypoint not found", http.StatusNotFound)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&waypoint); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := database.DB.Model(&waypoint).Updates(waypoint).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(waypoint); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Waypoint delete-request handler
func deleteWaypoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var waypoint structs.Waypoint
	if err := database.DB.First(&waypoint, params["id"]).Error; err != nil {
		http.Error(w, "Error 404: Waypoint not found", http.StatusNotFound)
		return
	}
	if err := database.DB.Delete(&waypoint).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
