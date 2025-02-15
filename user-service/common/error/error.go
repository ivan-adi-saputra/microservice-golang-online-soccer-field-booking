package error

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type ValidationResponse struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
}

var ErrValidator = map[string]string{}

func ErrValidationResponse(err error) (validationResponse []ValidationResponse) {
	var fieldErrors validator.ValidationErrors

	if errors.As(err, &fieldErrors) {
		for _, e := range fieldErrors {
			switch e.Tag() {
			case "required":
				validationResponse = append(validationResponse, ValidationResponse{
					Field:   e.Field(),
					Message: fmt.Sprintf("%s is required", e.Field()),
				})
			case "email":
				validationResponse = append(validationResponse, ValidationResponse{
					Field:   e.Field(),
					Message: fmt.Sprintf("%s is a valid email address", e.Field()),
				})
			default:
				if errValidator, ok := ErrValidator[e.Tag()]; ok {
					if count := strings.Count(errValidator, "%s"); count == 1 {
						validationResponse = append(validationResponse, ValidationResponse{
							Field:   e.Field(),
							Message: fmt.Sprintf(errValidator, e.Field()),
						})
					} else {
						validationResponse = append(validationResponse, ValidationResponse{
							Field:   e.Field(),
							Message: fmt.Sprintf(errValidator, e.Field(), e.Param()),
						})
					}
				} else {
					validationResponse = append(validationResponse, ValidationResponse{
						Field:   e.Field(),
						Message: fmt.Sprintf("something wrong on %s: %s", e.Field(), e.Tag()),
					})
				}
			}
		}
	}

	return validationResponse
}

func WrapError(err error) error {
	logrus.Errorf("error: %v", err)
	return err
}
