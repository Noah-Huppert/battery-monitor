package models

// User holds information about a person using the battery monitor server
type User struct {
	gorm.Model

	// FirstName holds a person's given name
	FirstName string

	// LastName holds the person's family name
	LastName string

	// Email holds the person's email
	Email string

	// EmailVerified designates if a user's email can be trusted to belong
	// to them
	EmailVerified bool

	// Devices will be populated with all the devices a user owns
	Devices []Devices
}
