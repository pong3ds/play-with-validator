package main

import (
	"fmt"
)

// CustomError is the custom error
type CustomError interface {
	Error() string
	GetStatus() int
	JSON() StandardError
}

// StandardError is the standard error
type StandardError struct {
	Status  int         `json:"-"`
	Code    string      `json:"code"`
	Message interface{} `json:"message"`
}

// JSON return JSON
func (c StandardError) JSON() StandardError {
	return c
}

// Error return error message
func (c StandardError) Error() string {
	return fmt.Sprintf("code : %v message : %v", c.Code, c.Message)
}

// GetStatus return status
func (c StandardError) GetStatus() int {
	return c.Status
}

// NewError return StandardError
func NewError(err error, errorType StandardError) StandardError {
	//TODO : CALL LOG SERVER OR SOMETHING
	fmt.Println("[Error]: ", err.Error())
	return errorType
}

// Recover is the recover func
func Recover(textError string) {
	if r := recover(); r != nil {
		panic(textError)
	}
}
