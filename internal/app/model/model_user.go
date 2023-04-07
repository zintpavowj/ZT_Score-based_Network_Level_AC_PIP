package model

import (
	"fmt"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// User ...
type User struct {
	ID                     int       `json:"id"`
	Name                   string    `json:"name"`
	Email                  string    `json:"email"`
	LastAccessTime         time.Time `json:"last_access_time"`
	Expected               float32   `json:"expected"`
	AccessTimeMin          time.Time `json:"access_time_min"`
	AccessTimeMax          time.Time `json:"access_time_max"`
	DatabaseUpdateTime     time.Time `json:"database_update_time"`
	PasswordFailedAttempts int       `json:"password_failed_attempts"`
}

// Validate ...
func (u *User) Validate() error {

	fmt.Println("ToDo: add correct user validation!")

	return validation.ValidateStruct(
		u,
		validation.Field(
			&u.Name,
			validation.Required,
			validation.Length(2, 50),
		),
		validation.Field(
			&u.Email,
			validation.Required,
			is.Email,
		),
		validation.Field(
			&u.DatabaseUpdateTime,
			validation.Date("2011-08-12 20:17:46 +1:00"),
		),
	)
}
