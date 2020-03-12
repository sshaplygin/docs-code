package go_codes_validator

import "errors"

var (
	ErrInvalidINNLength    = errors.New("invalid inn length")
	ErrInvalidBIKLength    = errors.New("invalid bik length")
	ErrInvalidKPPLength    = errors.New("invalid kpp length")
	ErrInvalidOGRNLength   = errors.New("invalid ogrn length")
	ErrInvalidOGRNIPLength = errors.New("invalid ogrinp length")
	ErrInvalidSNILSLength  = errors.New("invalid snils length")

	ErrInvalidFormattedSNILSLength = errors.New("invalid formatted snils length")

	ErrInvalidValue = errors.New("invalid code value")

	ErrInvalidBIKCountryCode = errors.New("invalid bik country code")
)
