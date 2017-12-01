package models

// Device holds information about phones user's own
type Device struct {
	gorm.Model

	// Name of device
	Name string

	// UserID identifies which user the device belongs to
	UserID uint

	// User is populated with the user who owns the device
	User User
}
