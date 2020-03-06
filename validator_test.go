package go_codes_validator

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateBIK(t *testing.T) {
	t.Parallel()

	t.Run("", func(t *testing.T) {

	})
}

type TestINNCase struct {
	INN     string
	IsValid bool
	Error   error
}

func TestValidateINN(t *testing.T) {
	t.Parallel()

	t.Run("invalid inn length", func(t *testing.T) {
		testCases := []TestINNCase{
			TestINNCase{
				INN:     "12345678",
				Error:   errInvalidINNLength,
				IsValid: false,
			},
			TestINNCase{
				INN:     "9876543211123",
				Error:   errInvalidINNLength,
				IsValid: false,
			},
			TestINNCase{
				INN:     "7707083893",
				Error:   nil,
				IsValid: true,
			},
			TestINNCase{
				INN:     "526317984689",
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
		testCases := []TestINNCase{
			TestINNCase{
				INN:     "77$7083893",
				Error:   errInvalidValue,
				IsValid: false,
			},
			TestINNCase{
				INN:     "98754321N123",
				Error:   errInvalidValue,
				IsValid: false,
			},
			TestINNCase{
				INN:     "9854132d1123",
				Error:   errInvalidValue,
				IsValid: false,
			},
			TestINNCase{
				INN:     "7707083893",
				Error:   nil,
				IsValid: true,
			},
			TestINNCase{
				INN:     "526317984689",
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
