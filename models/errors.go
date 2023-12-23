package models

import (
	"errors"
	"fmt"
)

var (
	// ErrInvalidValue invalid input value
	ErrInvalidValue = errors.New("invalid code value")

	// ErrInvalidLength invalid input document code length
	ErrInvalidLength = errors.New("invalid length")
)

// CommonError common error wrapped base error
type CommonError struct {
	Method string
	Err    error
}

func (c *CommonError) Error() string {
	return fmt.Sprintf("%s: %s", c.Method, c.Err.Error())
}
