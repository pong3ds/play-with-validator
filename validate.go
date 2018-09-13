package main

import (
	"strings"

	"github.com/3dsinteractive/govalidator"
)

// CustomValidator is custom validator
type CustomValidator struct{}

type jsonErr struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// FieldError is field error
type FieldError struct {
	Code   string      `json:"code"`
	Fields interface{} `json:"fields"`
}

func (f FieldError) Error() string {
	return "Code : " + f.Code
}

// ErrorToJson is error to json
func ErrorToJson(err error) (m map[string]jsonErr) {
	m = make(map[string]jsonErr)
	for _, value := range err.(govalidator.Errors) {
		m[value.(govalidator.Error).Name] = jsonErr{
			Code:    strings.ToUpper(value.(govalidator.Error).Validator),
			Message: value.(govalidator.Error).Err.Error(),
		}
	}
	return
}

// Validate is validate function
func (cv *CustomValidator) Validate(i interface{}) error {
	govalidator.SetFieldsRequiredByDefault(true)
	defer Recover("Validator has errors")

	if _, err := govalidator.ValidateStruct(i); err != nil {
		return FieldError{
			Code:   "INVALID_PARAMS",
			Fields: ErrorToJson(err),
		}
	}

	return nil
}
