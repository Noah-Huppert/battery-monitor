package models

// User holds information about a person using the battery monitor server
type User struct {
	Base

	// FirstName holds a person's given name
	FirstName string `gorm:"not null"`

	// LastName holds the person's family name
	LastName string `gorm:"not null"`

	// Email holds the person's email
	Email string `gorm:"not null;unique_index"`

	// EmailVerified designates if a user's email can be trusted to belong
	// to them
	EmailVerified bool `gorm:"not null"`

	// Devices will be populated with all the devices a user owns
	Devices []Devices `gorm:"ForeignKey:UserID"`
}
