package kpp

import "errors"

var (
	// ErrNilKPP try call methods for nil kpp struct
	ErrNilKPP = errors.New("nil kpp struct")

	// ErrRegistrationReasonCode invalid registration reason code
	ErrRegistrationReasonCode = errors.New("invalid registration reason code")

	// ErrInvalidSubjectCode
	ErrInvalidSubjectCode = errors.New("")

	// ErrInvalidRecorderCode
	ErrInvalidRecorderCode = errors.New("")

	// ErrInvalidReasonCode
	ErrInvalidReasonCode = errors.New("")

	// ErrInvalidSerialNumber
	ErrInvalidSerialNumber = errors.New("")
)
