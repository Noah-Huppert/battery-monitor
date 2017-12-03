package models

type AuthToken struct {
	Base

	UserID   uint
	DeviceID uint
	IssuedAt time.Time
}
