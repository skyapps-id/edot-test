package validator

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func SetupValidator() *validator.Validate {
	v := validator.New()

	// Register your custom validator here
	// https://pkg.go.dev/github.com/go-playground/validator/v10#hdr-Custom_Validation_Functions
	// v.RegisterValidation("name of your custom tag", validationFunc)

	return v
}

func Validate(c echo.Context, s interface{}) (err error) {
	if err = c.Bind(s); err != nil {
		err = fmt.Errorf("error bind : %s", err.Error())
		return
	}

	if err = c.Validate(s); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			err = formatValidationError(errs, s)
			return
		}
		err = fmt.Errorf("error validate : %s", err.Error())
	}

	return
}

func formatValidationError(errs validator.ValidationErrors, s interface{}) error {
	var sb strings.Builder

	structType := reflect.TypeOf(s)
	if structType.Kind() == reflect.Ptr {
		structType = structType.Elem()
	}

	for _, e := range errs {
		jsonTag := getJSONFieldName(structType, e.StructField())

		tag := e.Tag()
		var msg string

		switch tag {
		case "required":
			msg = fmt.Sprintf("The field '%s' is required", jsonTag)
		case "min":
			msg = fmt.Sprintf("The field '%s' must have at least %s item(s)", jsonTag, e.Param())
		case "max":
			msg = fmt.Sprintf("The field '%s' must have at most %s item(s)", jsonTag, e.Param())
		case "email":
			msg = fmt.Sprintf("The field '%s' must be a valid email address", jsonTag)
		case "gtfield":
			otherField := e.Param()
			msg = fmt.Sprintf("The field '%s' must be greater than '%s'", jsonTag, otherField)
		case "oneof":
			msg = fmt.Sprintf("The field '%s' must be one of [%s]", jsonTag, e.Param())
		default:
			msg = fmt.Sprintf("The field '%s' is invalid", jsonTag)
		}

		sb.WriteString(msg + ". ")
	}

	return errors.New(sb.String())
}

func getJSONFieldName(structType reflect.Type, fieldName string) string {
	if field, ok := structType.FieldByName(fieldName); ok {
		jsonTag := field.Tag.Get("json")
		if jsonTag != "" && jsonTag != "-" {
			return strings.Split(jsonTag, ",")[0]
		}
	}

	return fieldName
}
