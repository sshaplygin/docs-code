package snils

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	ru_doc_code "github.com/mrfoe7/go-codes-validator"
)

func TestValidate(t *testing.T) {
	t.Parallel()

	t.Run("invalid snils length", func(t *testing.T) {
		testCases := []ru_doc_code.TestCodeCase{
			{
				Code:    "112-233-445 95",
				Error:   nil,
				IsValid: true,
			},
			{
				Code:    "646-663-083 23",
				Error:   nil,
				IsValid: true,
			},
			{
				Code:    "112-233-445 951213",
				Error:   ru_doc_code.ErrInvalidSNILSLength,
				IsValid: false,
			},
			{
				Code:    "112-233 95",
				Error:   ru_doc_code.ErrInvalidSNILSLength,
				IsValid: false,
			},
		}
		for _, test := range testCases {
			isValid, err := Validate(test.Code)
			assert.Equal(t, test.IsValid, isValid, test.Code)
			assert.Equal(t, true, errors.Is(test.Error, err), test.Code)
		}
	})

	t.Run("invalid snils value", func(t *testing.T) {
		testCases := []ru_doc_code.TestCodeCase{
			{
				Code:    "112-233?445 95",
				Error:   ru_doc_code.ErrInvalidFormattedSNILSLength,
				IsValid: false,
			},
			{
				Code:    "1M2-234-445 95",
				Error:   ru_doc_code.ErrInvalidValue,
				IsValid: false,
			},
			{
				Code:    "112-233-445 98",
				Error:   nil,
				IsValid: false,
			},
			{
				Code:    "112-233-445#95",
				Error:   ru_doc_code.ErrInvalidFormattedSNILSLength,
				IsValid: false,
			},
			{
				Code:    "112-233-445 95",
				Error:   nil,
				IsValid: true,
			},
			{
				Code:    "646-663-083 23",
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
