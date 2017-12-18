package models

import "fmt"

// Auth Token stores details about HTTP API keys issued to users
type AuthToken struct {
	Base

	// UserID is the ID of the user the auth token was issued to
	UserID uint `gorm:"index"`

	// User will be populated with the user which the auth token was
	// issued to
	User User `gorm:"ForeignKey:UserID"`

	// DeviceID is the ID the device the auth token was issued to
	DeviceID uint `gorm:"index"`

	// Device will be populated with the device which the auth token was
	// issued to
	Device Device `gorm:"ForeignKey:DeviceID"`

	// Revoked specifies if the auth token is revoked
	Revoked bool `gorm:"not null"`
}

// Validate ensures the all model fields are not empty
func (t AuthToken) Validate() error {
	// Holds names of empty fields
	empty := []string{}

	// Check user id
	if len(t.UserID) == 0 {
		empty = append(empty, "UserID")
	}

	// Check device id
	if len(t.DeviceID) == 0 {
		empty = append(empty, "DeviceID")
	}

	// Check if any empty fields
	if len(empty) != 0 {
		return fmt.Errorf("the following fields were empty: %s",
			Strings.join(empty))
	}

	// All good
	return nil
}
