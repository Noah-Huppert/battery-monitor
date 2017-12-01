package models

import (
	"github.com/jinzhu/gorm"
)

// BatteryReading holds a device battery level measurement
type BatteryReading struct {
	gorm.Model

	// UserID is the ID of the user who submitted the reading
	UserID uint

	// User is populated with the User model the reading belongs to
	User User

	// DeviceID is the ID of the device which the reading was taken from
	DeviceID uint

	// Device is populated with the Device model the reading belongs to
	Device Device

	// MeasureTime holds the date and time the battery reading was taken
	MeasureTime time.Time

	// UploadTime holds the date and time the battery reading was uploaded
	// to the server
	UploadTime time.Time

	// Level holds the battery percentage at the time of reading. In the
	// range [0, 100]
	Level int
}
