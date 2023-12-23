package bik

import "errors"

var (
	// ErrInvalidCountryCode invalid bik code country
	ErrInvalidCountryCode = errors.New("invalid bik country code")
)
