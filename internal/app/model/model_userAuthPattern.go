package model

import validation "github.com/go-ozzo/ozzo-validation"

// UserAuthPattern ...
type UserAuthPattern struct {
	ID                  int    `json:"id"`
	UserAuthPatternName string `json:"name"`
}

// Validate ...
func (uap *UserAuthPattern) Validate() error {
	return validation.ValidateStruct(
		uap,
		validation.Field(
			&uap.UserAuthPatternName,
			validation.Required,
			validation.Length(2, 50),
		),
	)
}
