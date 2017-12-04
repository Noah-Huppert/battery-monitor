package models

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
