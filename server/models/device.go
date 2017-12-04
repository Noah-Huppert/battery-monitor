package models

// Device holds information about phones a user owns
type Device struct {
	Base

	// UserID identifies which user the device belongs to
	UserID uint `gorm:"index"`

	// User is populated with the user who owns the device
	User User `gorm:"ForeignKey:UserID"`

	// Name of device
	Name string `gorm:"not null"`

	// Fingerprint is a unique value used to identify the device
	Fingerprint string `gorm:"not null"`
}
