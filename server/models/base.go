package models

import "time"

// Base is the struct which contains the default fields which all models stored
// in the database have
type Base struct {
	// ID is the unique identifier of the model
	ID uint `gorm:"primary_key;AUTO_INCREMENT"`

	// CreatedAt specifies when the model was initially saved in the db
	CreatedAt time.Time `gorm:"not null"`

	// UpdatedAt specified when the model was last updated
	UpdatedAt time.Time `gorm:"not null"`
}
