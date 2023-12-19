package bik

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	ru_doc_code "github.com/sshaplygin/ru-doc-code"
)

func TestValidate(t *testing.T) {
	t.Parallel()

	t.Run("invalid bik length", func(t *testing.T) {
		testCases := []ru_doc_code.TestCodeCase{
			{
				Code:    "1234567888776",
				Error:   ru_doc_code.ErrInvalidLength,
				IsValid: false,
			},
			{
				Code:    "044525",
				Error:   ru_doc_code.ErrInvalidLength,
				IsValid: false,
			},
			{
				Code:    "044525225",
				Error:   nil,
				IsValid: true,
			},
			{
				Code:    "044525012",
				Error:   nil,
				IsValid: true,
			},
		}
		for i, test := range testCases {
			isValid, err := Validate(test.Code)
			assert.Equal(t, test.IsValid, isValid, test.Code)
			if err != nil {
				assert.ErrorAs(t, err, &test.Error, fmt.Sprintf("invalid test case %d: input: %s", i, test.Code))
			} else {
				assert.NoError(t, err, fmt.Sprintf("invalid test case %d: input: %s", i, test.Code))
			}
		}
	})

	t.Run("invalid bik value", func(t *testing.T) {
		testCases := []ru_doc_code.TestCodeCase{
			{
				Code:    "0445?5226",
				Error:   ru_doc_code.ErrInvalidValue,
				IsValid: false,
			},
			{
				Code:    "054525225",
				Error:   ru_doc_code.ErrInvalidBIKCountryCode,
				IsValid: false,
			},
			{
				Code:    "104525225",
				Error:   ru_doc_code.ErrInvalidBIKCountryCode,
				IsValid: false,
			},
			{
				Code:    "044#55#25",
				Error:   ru_doc_code.ErrInvalidValue,
				IsValid: false,
			},
			{
				Code:    "044525225",
				Error:   nil,
				IsValid: true,
			},
			{
				Code:    "044525012",
				Error:   nil,
				IsValid: true,
			},
		}
		for i, test := range testCases {
			isValid, err := Validate(test.Code)
			assert.Equal(t, test.IsValid, isValid, test.Code, test.IsValid)
			if err != nil {
				assert.ErrorAs(t, err, &test.Error, fmt.Sprintf("invalid test case %d: input: %s", i, test.Code))
			} else {
				assert.Empty(t, err, fmt.Sprintf("invalid test case %d: input: %s", i, test.Code))
			}
		}
	})
}
