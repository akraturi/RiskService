package server

import "github.com/go-playground/validator/v10"

func NewRequestValidator() *validator.Validate {
	validate := validator.New()
	err := validate.RegisterValidation("oneof", RiskStateValidator)
	if err != nil {
		return nil
	}
	return validate
}
