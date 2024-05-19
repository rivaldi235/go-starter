package validation

import (
	"service-code/model/dto/json"
	"strings"
	"unicode"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"github.com/stoewer/go-strcase"
)

func GetValidationError(err error) []json.ValidationField {
	var validationFields []json.ValidationField
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, validationError := range ve {
			log.Debug().Msgf("validationError: %v", validationError)
			myField := convertFieldRequired(validationError.Namespace())
			validationFields = append(validationFields, json.ValidationField{
				FieldName: myField,
				Message:   formatMessage(validationError),
			})
		}
	}
	return validationFields
}

func convertFieldRequired(myValue string) string {
	log.Debug().Msg("convertFieldRequired: " + myValue)
	fieldSegments := strings.Split(myValue, ".")
	var myField string
	for i, val := range fieldSegments {
		if i == 0 {
			continue
		}
		if i == len(fieldSegments)-1 {
			myField += strcase.SnakeCase(val)
		} else {
			myField += strcase.LowerCamelCase(val)
		}
	}
	return myField
}

func formatMessage(err validator.FieldError) string {
	var message string

	switch err.Tag() {
	case "required":
		message = "required"
	case "number":
		message = "must be number"
	case "email":
		message = "invalid format email"
	case "DateOnly":
		message = "invalid format date"
	case "min":
		message = "minimum value is not exceed"
	case "max":
		message = "max value is exceed"
	}

	return message
}

func ValidatePasswordFormat(password string) bool {
	if len(password) < 8 {
		return false
	}

	var (
		hasUpper   bool
		hasLower   bool
		hasDigit   bool
		hasSpecial bool
	)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasDigit = true
		case strings.ContainsRune("!@#$%^&*()-_+=[]{};:'\"<>,.?/~`\\|", char):
			hasSpecial = true
		}
	}

	return hasUpper && hasLower && hasDigit && hasSpecial
}
