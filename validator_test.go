package ru_doc_code

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
			assert.Equal(t, test.IsValid, isValid, test.Code)
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
			assert.Equal(t, test.IsValid, isValid, test.Code, test.IsValid)
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
			assert.Equal(t, test.IsValid, isValid, test.Code)
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
			assert.Equal(t, test.IsValid, isValid, test.Code)
			assert.Equal(t, true, errors.Is(test.Error, err), test.Code)
		}
	})
}

func TestValidateOGRN(t *testing.T) {
	t.Parallel()

	t.Run("invalid ogrn length", func(t *testing.T) {
		testCases := []TestCodeCase{
			TestCodeCase{
				Code:    "1027700132195",
				Error:   nil,
				IsValid: true,
			},
			TestCodeCase{
				Code:    "1027739244741",
				Error:   nil,
				IsValid: true,
			},
			TestCodeCase{
				Code:    "102773924",
				Error:   ErrInvalidOGRNLength,
				IsValid: false,
			},
			TestCodeCase{
				Code:    "10277392447411231",
				Error:   ErrInvalidOGRNLength,
				IsValid: false,
			},
		}
		for _, test := range testCases {
			isValid, err := IsOGRNValid(test.Code)
			assert.Equal(t, test.IsValid, isValid, test.Code)
			assert.Equal(t, true, errors.Is(test.Error, err), test.Code)
		}
	})

	t.Run("invalid ogrn value", func(t *testing.T) {
		testCases := []TestCodeCase{
			TestCodeCase{
				Code:    "102773??44741",
				Error:   ErrInvalidValue,
				IsValid: false,
			},
			TestCodeCase{
				Code:    "1027739244742",
				Error:   nil,
				IsValid: false,
			},
			TestCodeCase{
				Code:    "10@7739244%42",
				Error:   ErrInvalidValue,
				IsValid: false,
			},
			TestCodeCase{
				Code:    "1027700132195",
				Error:   nil,
				IsValid: true,
			},
			TestCodeCase{
				Code:    "1027739244741",
				Error:   nil,
				IsValid: true,
			},
		}
		for _, test := range testCases {
			isValid, err := IsOGRNValid(test.Code)
			assert.Equal(t, test.IsValid, isValid, test.Code, test.IsValid)
			assert.Equal(t, true, errors.Is(test.Error, err), test.Code)
		}
	})
}

func TestValidateOGRNIP(t *testing.T) {
	t.Parallel()

	t.Run("invalid ogrnip length", func(t *testing.T) {
		testCases := []TestCodeCase{
			TestCodeCase{
				Code:    "304500116000157",
				Error:   nil,
				IsValid: true,
			},
			TestCodeCase{
				Code:    "312502904600034",
				Error:   nil,
				IsValid: true,
			},
			TestCodeCase{
				Code:    "31250290460",
				Error:   ErrInvalidOGRNIPLength,
				IsValid: false,
			},
			TestCodeCase{
				Code:    "3045001111236000157",
				Error:   ErrInvalidOGRNIPLength,
				IsValid: false,
			},
		}
		for _, test := range testCases {
			isValid, err := IsOGRNIPValid(test.Code)
			assert.Equal(t, test.IsValid, isValid, test.Code)
			assert.Equal(t, true, errors.Is(test.Error, err), test.Code)
		}
	})

	t.Run("invalid ogrnip value", func(t *testing.T) {
		testCases := []TestCodeCase{
			TestCodeCase{
				Code:    "312502??4600034",
				Error:   ErrInvalidValue,
				IsValid: false,
			},
			TestCodeCase{
				Code:    "304500116000158",
				Error:   nil,
				IsValid: false,
			},
			TestCodeCase{
				Code:    "512502904600034",
				Error:   ErrInvalidValue,
				IsValid: false,
			},
			TestCodeCase{
				Code:    "304500116000157",
				Error:   nil,
				IsValid: true,
			},
			TestCodeCase{
				Code:    "312502904600034",
				Error:   nil,
				IsValid: true,
			},
		}
		for _, test := range testCases {
			isValid, err := IsOGRNIPValid(test.Code)
			assert.Equal(t, test.IsValid, isValid, test.Code, test.IsValid)
			assert.Equal(t, true, errors.Is(test.Error, err), test.Code)
		}
	})
}

func TestValidateSNILS(t *testing.T) {
	t.Parallel()

	t.Run("invalid snils length", func(t *testing.T) {
		testCases := []TestCodeCase{
			TestCodeCase{
				Code:    "112-233-445 95",
				Error:   nil,
				IsValid: true,
			},
			TestCodeCase{
				Code:    "646-663-083 23",
				Error:   nil,
				IsValid: true,
			},
			TestCodeCase{
				Code:    "112-233-445 951213",
				Error:   ErrInvalidSNILSLength,
				IsValid: false,
			},
			TestCodeCase{
				Code:    "112-233 95",
				Error:   ErrInvalidSNILSLength,
				IsValid: false,
			},
		}
		for _, test := range testCases {
			isValid, err := IsSNILSValid(test.Code)
			assert.Equal(t, test.IsValid, isValid, test.Code)
			assert.Equal(t, true, errors.Is(test.Error, err), test.Code)
		}
	})

	t.Run("invalid snils value", func(t *testing.T) {
		testCases := []TestCodeCase{
			TestCodeCase{
				Code:    "112-233?445 95",
				Error:   ErrInvalidFormattedSNILSLength,
				IsValid: false,
			},
			TestCodeCase{
				Code:    "1M2-234-445 95",
				Error:   ErrInvalidValue,
				IsValid: false,
			},
			TestCodeCase{
				Code:    "112-233-445 98",
				Error:   nil,
				IsValid: false,
			},
			TestCodeCase{
				Code:    "112-233-445#95",
				Error:   ErrInvalidFormattedSNILSLength,
				IsValid: false,
			},
			TestCodeCase{
				Code:    "112-233-445 95",
				Error:   nil,
				IsValid: true,
			},
			TestCodeCase{
				Code:    "646-663-083 23",
				Error:   nil,
				IsValid: true,
			},
		}
		for _, test := range testCases {
			isValid, err := IsSNILSValid(test.Code)
			assert.Equal(t, test.IsValid, isValid, test.Code, test.IsValid)
			assert.Equal(t, true, errors.Is(test.Error, err), test.Code)
		}
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

	t.Run("invalid kpp value", func(t *testing.T) {
		testCases := []TestCodeCase{
			TestCodeCase{
				Code:    "773$N3301",
				Error:   ErrInvalidValue,
				IsValid: false,
			},
			TestCodeCase{
				Code:    "7736#3&01",
				Error:   ErrInvalidValue,
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
