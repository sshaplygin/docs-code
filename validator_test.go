package go_codes_validator

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCodeCase struct {
	Code    string
	IsValid bool
	Error   error
}

func TestValidateBIK(t *testing.T) {
	t.Parallel()

	t.Run("", func(t *testing.T) {

	})
}

func TestValidateINN(t *testing.T) {
	t.Parallel()

	t.Run("invalid inn length", func(t *testing.T) {
		testCases := []TestCodeCase{
			TestCodeCase{
				Code:    "12345678",
				Error:   errInvalidINNLength,
				IsValid: false,
			},
			TestCodeCase{
				Code:    "9876543211123",
				Error:   errInvalidINNLength,
				IsValid: false,
			},
			TestCodeCase{
				Code:    "7707083893",
				Error:   nil,
				IsValid: true,
			},
			TestCodeCase{
				Code:    "526317984689",
				Error:   nil,
				IsValid: true,
			},
		}
		for _, test := range testCases {
			isValid, err := ValidateINN(test.INN)
			assert.Equal(t, isValid, test.IsValid)
			assert.Equal(t, errors.Is(test.Error, err), true)
		}
	})

	t.Run("invalid inn value", func(t *testing.T) {
		testCases := []TestCodeCase{
			TestCodeCase{
				Code:    "77$7083893",
				Error:   errInvalidValue,
				IsValid: false,
			},
			TestCodeCase{
				Code:    "98754321N123",
				Error:   errInvalidValue,
				IsValid: false,
			},
			TestCodeCase{
				Code:    "9854132d1123",
				Error:   errInvalidValue,
				IsValid: false,
			},
			TestCodeCase{
				Code:    "7707083893",
				Error:   nil,
				IsValid: true,
			},
			TestCodeCase{
				Code:    "526317984689",
				Error:   nil,
				IsValid: true,
			},
		}
		for _, test := range testCases {
			isValid, err := ValidateINN(test.Code)
			assert.Equal(t, isValid, test.IsValid)
			assert.Equal(t, errors.Is(test.Error, err), true)
		}
	})
}

func TestValidateOGRN(t *testing.T) {
	t.Parallel()

	t.Run("", func(t *testing.T) {

	})
}

func TestValidateOGRNIP(t *testing.T) {
	t.Parallel()

	t.Run("", func(t *testing.T) {

	})
}

func TestValidateSNILS(t *testing.T) {
	t.Parallel()

	t.Run("", func(t *testing.T) {

	})
}

func TestValidateKPP(t *testing.T) {
	t.Parallel()

	t.Run("", func(t *testing.T) {

	})
}
