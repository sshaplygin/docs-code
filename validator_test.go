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

func TestIsBIKValid(t *testing.T) {
	t.Parallel()

	t.Run("invalid bik length", func(t *testing.T) {
		testCases := []TestCodeCase{
			TestCodeCase{
				Code:    "1234567888776",
				Error:   ErrInvalidBIKLength,
				IsValid: false,
			},
			TestCodeCase{
				Code:    "044525",
				Error:   ErrInvalidBIKLength,
				IsValid: false,
			},
			TestCodeCase{
				Code:    "044525225",
				Error:   nil,
				IsValid: true,
			},
			TestCodeCase{
				Code:    "044525012",
				Error:   nil,
				IsValid: true,
			},
		}
		for _, test := range testCases {
			isValid, err := IsBIKValid(test.Code)
			assert.Equal(t, isValid, test.IsValid, test.Code)
			assert.Equal(t, true, errors.Is(test.Error, err), test.Code)
		}
	})

	t.Run("invalid bik value", func(t *testing.T) {
		testCases := []TestCodeCase{
			TestCodeCase{
				Code:    "0445?5226",
				Error:   ErrInvalidValue,
				IsValid: false,
			},
			TestCodeCase{
				Code:    "054525225",
				Error:   ErrInvalidBIKCountryCode,
				IsValid: false,
			},
			TestCodeCase{
				Code:    "104525225",
				Error:   ErrInvalidBIKCountryCode,
				IsValid: false,
			},
			TestCodeCase{
				Code:    "044#55#25",
				Error:   ErrInvalidValue,
				IsValid: false,
			},
			TestCodeCase{
				Code:    "044525225",
				Error:   nil,
				IsValid: true,
			},
			TestCodeCase{
				Code:    "044525012",
				Error:   nil,
				IsValid: true,
			},
		}
		for _, test := range testCases {
			isValid, err := IsBIKValid(test.Code)
			assert.Equal(t, isValid, test.IsValid, test.Code, test.IsValid)
			assert.Equal(t, true, errors.Is(test.Error, err), test.Code)
		}
	})
}

func TestIsINNValid(t *testing.T) {
	t.Parallel()

	t.Run("invalid inn length", func(t *testing.T) {
		testCases := []TestCodeCase{
			TestCodeCase{
				Code:    "12345678",
				Error:   ErrInvalidINNLength,
				IsValid: false,
			},
			TestCodeCase{
				Code:    "9876543211123",
				Error:   ErrInvalidINNLength,
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
			isValid, err := IsINNValid(test.Code)
			assert.Equal(t, isValid, test.IsValid, test.Code)
			assert.Equal(t, true, errors.Is(test.Error, err), test.Code)
		}
	})

	t.Run("invalid inn value", func(t *testing.T) {
		testCases := []TestCodeCase{
			TestCodeCase{
				Code:    "77$7083893",
				Error:   ErrInvalidValue,
				IsValid: false,
			},
			TestCodeCase{
				Code:    "98754321N123",
				Error:   ErrInvalidValue,
				IsValid: false,
			},
			TestCodeCase{
				Code:    "9854132d1123",
				Error:   ErrInvalidValue,
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
			isValid, err := IsINNValid(test.Code)
			assert.Equal(t, isValid, test.IsValid, test.Code)
			assert.Equal(t, true, errors.Is(test.Error, err), test.Code)
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

	t.Run("invalid kpp length", func(t *testing.T) {
		testCases := []TestCodeCase{
			TestCodeCase{
				Code:    "1234567888776",
				Error:   ErrInvalidKPPLength,
				IsValid: false,
			},
			TestCodeCase{
				Code:    "044525",
				Error:   ErrInvalidKPPLength,
				IsValid: false,
			},
			TestCodeCase{
				Code:    "773643301",
				Error:   nil,
				IsValid: true,
			},
			TestCodeCase{
				Code:    "773643001",
				Error:   nil,
				IsValid: true,
			},
		}
		for _, test := range testCases {
			isValid, err := IsKPPValid(test.Code)
			assert.Equal(t, isValid, test.IsValid, test.Code)
			assert.Equal(t, true, errors.Is(test.Error, err), test.Code)
		}
	})
}
