package validation

import (
	"api/internal/handler/httperr"
	"fmt"
	"reflect"
	"strings"
	"unicode"

	"github.com/go-playground/validator/v10"
)

var val *validator.Validate

func init() {
	// Cria o validador global e registra a validação personalizada
	val = validator.New(validator.WithRequiredStructEnabled())

	// Registro da função para obter o nome do campo JSON
	val.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	// Registra a validação personalizada para `password`
	val.RegisterValidation("password", validatePassword)
}

// Função personalizada para validar senhas
func validatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	var hasMinLen, hasUpper, hasLower, hasNumber, hasSpecial bool

	if len(password) >= 8 {
		hasMinLen = true
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}

// Validação dos dados HTTP
func ValidateHttpData(d interface{}) *httperr.RestErr {
	if err := val.Struct(d); err != nil {
		var errosCauses []httperr.Fields

		for _, e := range err.(validator.ValidationErrors) {
			cause := httperr.Fields{}
			fieldName := e.Field()

			switch e.Tag() {
			case "required":
				cause.Message = fmt.Sprintf("%s is required", fieldName)
				cause.Field = fieldName
				cause.Value = e.Value()
			case "uuid4":
				cause.Message = fmt.Sprintf("%s is not a valid uuid", fieldName)
				cause.Field = fieldName
				cause.Value = e.Value()
			case "boolean":
				cause.Message = fmt.Sprintf("%s is not a valid boolean", fieldName)
				cause.Field = fieldName
				cause.Value = e.Value()
			case "min":
				cause.Message = fmt.Sprintf("%s must be greater than %s", fieldName, e.Param())
				cause.Field = fieldName
				cause.Value = e.Value()
			case "max":
				cause.Message = fmt.Sprintf("%s must be less than %s", fieldName, e.Param())
				cause.Field = fieldName
				cause.Value = e.Value()
			case "email":
				cause.Message = fmt.Sprintf("%s is not a valid email", fieldName)
				cause.Field = fieldName
				cause.Value = e.Value()
			case "containsany":
				cause.Message = fmt.Sprintf("%s must contain at least one of the following characters: !@#$%%*", fieldName)
				cause.Field = fieldName
				cause.Value = e.Value()
			case "password":
				cause.Message = fmt.Sprintf("%s must have at least 8 characters, an uppercase letter, a lowercase letter, a number, and a special character", fieldName)
				cause.Field = fieldName
				cause.Value = e.Value()
			default:
				cause.Message = "invalid field"
				cause.Field = fieldName
				cause.Value = e.Value()
			}

			errosCauses = append(errosCauses, cause)
		}

		return httperr.NewBadRequestValidationError("some fields are invalid", errosCauses)
	}

	return nil
}
