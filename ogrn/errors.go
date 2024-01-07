package ogrn

import "errors"

var (
	// ErrNilOGRN try call methods for nil ogrn struct
	ErrNilOGRN = errors.New("nil ogrn struct")

	// ErrInvalidCodeType invalid code type
	ErrInvalidCodeType = errors.New("invalid code type")

	// ErrInvalidYearsNumbers invalid years number code
	ErrInvalidYearsNumbers = errors.New("invalid years number code")

	// ErrInvalidRegion invalid region code
	ErrInvalidRegion = errors.New("invalid region code")

	// ErrInvalidSerialNumbers invalid serial numbers
	ErrInvalidSerialNumbers = errors.New("invalid serial numbers")
)
