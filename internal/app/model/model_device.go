package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

// Device ...
type Device struct {
	ID             int       `json:"id"`
	DeviceName     string    `json:"name"`
	DeviceCertCN   string    `json:"cert_cn"`
	LastAccessTime time.Time `json:"last_access_time"`
	Expected       float32   `json:"expected"`
}

// Validate ...
func (s *Device) Validate() error {
	return validation.ValidateStruct(
		s,
		validation.Field(
			&s.DeviceName,
			validation.Required,
			validation.Length(2, 50),
		),
		validation.Field(
			&s.DeviceCertCN,
			validation.Required,
			validation.Length(2, 100),
		),
	)
}
