package kpp

import "errors"

var (
	// ErrNilKPP try call methods for nil kpp struct
	ErrNilKPP = errors.New("nil kpp struct")

	// ErrInvalidTaxRegion invalid tax region code
	ErrInvalidTaxRegion = errors.New("invalid tax region code")

	// ErrInvalidReasonCode invalid reason code
	ErrInvalidReasonCode = errors.New("invalid reason code")

	// ErrInvalidSerialNumbers invalid serial number
	ErrInvalidSerialNumbers = errors.New("invalid serial number")
)
