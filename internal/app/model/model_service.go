package model

import validation "github.com/go-ozzo/ozzo-validation"

// Service ...
type Service struct {
	ID                 int    `json:"id"`
	ServiceName        string `json:"name"`
	ServiceSNI         string `json:"sni"`
	DataSensitivity    int    `json:"data_sensitivity"`
	SoftwarePatchLevel string `json:"software_patch_level"`
}

// Validate ...
func (s *Service) Validate() error {
	return validation.ValidateStruct(
		s,
		validation.Field(
			&s.ServiceName,
			validation.Required,
			validation.Length(2, 50),
		),
		validation.Field(
			&s.ServiceSNI,
			validation.Required,
			validation.Length(2, 100),
		),
	)
}
