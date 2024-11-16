package db

import "gorm.io/gorm"

// Workspace model -> has many Endpoints
type Workspace struct {
	gorm.Model
	ID        string     `gorm:"primaryKey"`
	Name      string     `gorm:"not null;unique"` // Ensures workspace names are unique
	Endpoints []Endpoint `gorm:"foreignKey:WorkspaceID;references:ID"`
}

// Endpoint model
type Endpoint struct {
	gorm.Model
	ID              string `gorm:"primaryKey"`
	WorkspaceID     string `gorm:"not null"`
	Name            string `gorm:"not null"`
	URL             string `gorm:"not null;unique"` // Ensures each URL is unique
	StartHour       int8   `gorm:"not null"`        // Hour of the day [1-24]
	MonitorInterval int8   `gorm:"not null"`        // Monitoring interval, constant-based ["quarter","half","three-quarter","one"]
	TimeZone        string `gorm:"not null"`
}
