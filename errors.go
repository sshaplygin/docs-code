package ru_doc_code

import "errors"

var (
	// ErrInvalidINNLength
	ErrInvalidINNLength = errors.New("invalid inn length")

	// ErrInvalidBIKLength
	ErrInvalidBIKLength = errors.New("invalid bik length")

	// ErrInvalidKPPLength
	ErrInvalidKPPLength = errors.New("invalid kpp length")

	// ErrInvalidOGRNLength
	ErrInvalidOGRNLength = errors.New("invalid ogrn length")

	// ErrInvalidOGRNIPLength
	ErrInvalidOGRNIPLength = errors.New("invalid ogrinp length")

	// ErrInvalidSNILSLength
	ErrInvalidSNILSLength = errors.New("invalid snils length")

	// ErrInvalidFormattedSNILSLength
	ErrInvalidFormattedSNILSLength = errors.New("invalid formatted snils length")

	// ErrInvalidRegistrationReasonCode
	ErrInvalidRegistrationReasonCode = errors.New("invalid registration reason code")

	// ErrInvalidValue
	ErrInvalidValue = errors.New("invalid code value")

	// ErrInvalidBIKCountryCode
	ErrInvalidBIKCountryCode = errors.New("invalid bik country code")

	// ErrNotImplemented
	ErrNotImplemented = errors.New("method does not implemented")
)
