package ru_doc_code

import (
	"errors"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	// ErrInvalidLength invalid input document code length
	ErrInvalidLength = errors.New("invalid length")

	// ErrInvalidFormattedSNILSLength
	ErrInvalidFormattedSNILSLength = errors.New("invalid formatted snils length")

	// ErrRegistrationReasonCode
	ErrRegistrationReasonCode = errors.New("invalid registration reason code")

	// ErrInvalidValue
	ErrInvalidValue = errors.New("invalid code value")

	// ErrInvalidBIKCountryCode
	ErrInvalidBIKCountryCode = errors.New("invalid bik country code")

	// ErrNotImplemented
	ErrNotImplemented = errors.New("method does not implemented")
)

type CommonError struct {
	Method string
	Err    error
}

func (c *CommonError) Error() string {
	return fmt.Sprintf("%s: %s", c.Method, c.Err.Error())
}

func GetPackageName() (string, error) {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		return "", errors.New("invalid runtime caller")
	}
	parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	pl := len(parts)

	pathArr := strings.Split(parts[pl-2], string(filepath.Separator))
	if len(pathArr) == 0 {
		return "", errors.New("invalid path length")
	}
	pkgName := pathArr[len(pathArr)-1]

	return pkgName, nil
}
