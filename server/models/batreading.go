package models

import "time"

// BatteryReading holds a device battery level measurement
type BatteryReading struct {
	Base

	// UserID is the ID of the user who submitted the reading
	UserID uint `gorm:"index"`

	// User is populated with the User model the reading belongs to
	User User `gorm:"ForeignKey:UserID"`

	// DeviceID is the ID of the device which the reading was taken from
	DeviceID uint `gorm:"index"`

	// Device is populated with the Device model the reading belongs to
	Device Device `gorm:"ForeignKey:DeviceID"`

	// MeasureTime holds the date and time the battery reading was taken
	MeasureTime time.Time `gorm:"not null"`

	// UploadTime holds the date and time the battery reading was uploaded
	// to the server
	UploadTime time.Time `gorm:"not null"`

	// Level holds the battery percentage at the time of reading. In the
	// range [0, 100]
	Level int `gorm:"not null"`
}
