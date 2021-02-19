package ogrnip

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	ru_doc_code "github.com/mrfoe7/go-codes-validator"
)

func TestValidate(t *testing.T) {
	t.Parallel()

	t.Run("invalid ogrnip length", func(t *testing.T) {
		testCases := []ru_doc_code.TestCodeCase{
			{
				Code:    "304500116000157",
				Error:   nil,
				IsValid: true,
			},
			{
				Code:    "312502904600034",
				Error:   nil,
				IsValid: true,
			},
			{
				Code:    "31250290460",
				Error:   ru_doc_code.ErrInvalidOGRNIPLength,
				IsValid: false,
			},
			{
				Code:    "3045001111236000157",
				Error:   ru_doc_code.ErrInvalidOGRNIPLength,
				IsValid: false,
			},
		}
		for _, test := range testCases {
			isValid, err := Validate(test.Code)
			assert.Equal(t, test.IsValid, isValid, test.Code)
			assert.Equal(t, true, errors.Is(test.Error, err), test.Code)
		}
	})

	t.Run("invalid ogrnip value", func(t *testing.T) {
		testCases := []ru_doc_code.TestCodeCase{
			{
				Code:    "312502??4600034",
				Error:   ru_doc_code.ErrInvalidValue,
				IsValid: false,
			},
			{
				Code:    "304500116000158",
				Error:   nil,
				IsValid: false,
			},
			{
				Code:    "512502904600034",
				Error:   ru_doc_code.ErrInvalidValue,
				IsValid: false,
			},
			{
				Code:    "304500116000157",
				Error:   nil,
				IsValid: true,
			},
			{
				Code:    "312502904600034",
				Error:   nil,
				IsValid: true,
			},
		}
		for _, test := range testCases {
			isValid, err := Validate(test.Code)
			assert.Equal(t, test.IsValid, isValid, test.Code, test.IsValid)
			assert.Equal(t, true, errors.Is(test.Error, err), test.Code)
		}
	})
}
