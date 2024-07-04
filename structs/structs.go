package structs

import "time"

// Client struct
type Client struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `json:"name"`
	Rides []Ride `gorm:"foreignKey:ClientID"`
}

// Ride struct
type Ride struct {
	ID        uint       `gorm:"primaryKey"`
	RideDate  time.Time  `json:"rideDate"`
	Distance  int        `json:"distance"`
	Price     int        `json:"price"`
	ClientID  uint       `json:"clientId"`
	DriverID  uint       `json:"driverId"`
	Waypoints []Waypoint `gorm:"foreignKey:RideID;constraint:OnDelete:CASCADE;"`
}

// Driver struct
type Driver struct {
	ID           uint   `gorm:"primaryKey"`
	Name         string `json:"name"`
	LicensePlate string `json:"licensePlate"`
	Rides        []Ride `gorm:"foreignKey:DriverID"`
}

// Waypoint struct
type Waypoint struct {
	ID        uint    `gorm:"primaryKey"`
	Number    int     `json:"number"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	RideID    uint    `gorm:"constraint:OnDelete:CASCADE;"`
}
