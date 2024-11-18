package database

import (
	"gorm.io/gorm"
)

// Workspace model -> has many Endpoints
type Workspace struct {
	gorm.Model
	ID          string `gorm:"primaryKey"`
	Name        string `gorm:"not null;unique"` // Ensures workspace names are unique
	Description string
	Endpoints   []Endpoint `gorm:"foreignKey:WorkspaceID;references:ID"`
}

type Endpoint struct {
	gorm.Model
	ID              string `gorm:"primaryKey"`
	WorkspaceID     string `gorm:"not null"`
	Name            string `gorm:"not null"`
	Description     string
	URL             string `gorm:"not null;unique"`
	LastMonitored   int8   `gorm:"not null"` // Hour of the day [1-24]
	MonitorInterval string `gorm:"not null"` // Use predefined constants
	ToMonitor       int8   `gorm:"not null"`
	LastStatus      string `gorm:"not null"` // Constant-based status [active, down, unknown]
	// TimeZone         string    `gorm:"not null"` (for now this shouldn't be an issue , execting in a same timezone)
}
