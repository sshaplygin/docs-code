package bik

import "errors"

var (
	// ErrNilBIK try call methods for nil bik struct
	ErrNilBIK = errors.New("nil bik struct")

	// ErrInvalidCountryCode invalid bik code country
	ErrInvalidCountryCode = errors.New("invalid bik country code")

	// ErrInvalidTerritoryCode invalid okato territorial code of the subject
	ErrInvalidTerritoryCode = errors.New("invalid okato city code")

	// ErrInvalidUnitConditionalNumber invalid unit conditional number
	ErrInvalidUnitConditionalNumber = errors.New("invalid unit conditional number")

	// ErrInvalidLastAccountNumbers invalid okato territorial code of the subject
	ErrInvalidLastAccountNumbers = errors.New("invalid last account number")
)
