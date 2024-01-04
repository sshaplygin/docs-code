package kpp

import "errors"

var (
	// ErrNilKPP try call methods for nil kpp struct
	ErrNilKPP = errors.New("nil kpp struct")

	// ErrRegistrationReasonCode invalid registration reason code
	ErrRegistrationReasonCode = errors.New("invalid registration reason code")
)
