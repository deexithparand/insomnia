package database

import (
	"time"

	"gorm.io/gorm"
)

// Workspace model -> has many Endpoints
type Workspace struct {
	gorm.Model
	Name        string     `gorm:"not null;unique"` // Unique workspace name
	Description *string    // Optional: NULLable description
	Endpoints   []Endpoint `gorm:"foreignKey:WorkspaceID;constraint:OnDelete:CASCADE;"` // Cascade delete
}

// Endpoint model
type Endpoint struct {
	gorm.Model
	WorkspaceID      uint      `gorm:"not null"` // References Workspace ID
	Name             string    `gorm:"not null"` // Endpoint name
	Description      *string   // Optional description
	URL              string    `gorm:"not null;uniqueIndex:workspace_url"` // Unique within a workspace
	StartHour        int8      `gorm:"not null"`                           // [1-24] (hour of the day)
	MonitorInterval  string    `gorm:"not null"`                           // Constant-based, e.g., "quarter", "half", "three-quarter", "whole"
	TimeZone         string    `gorm:"not null"`                           // Timezone of the endpoint
	LastResponseTime time.Time // Timestamp of the last response
	LastStatus       string    `gorm:"not null"` // Constant-based status [active, down, unknown]
}
