package ogrn

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	ru_doc_code "github.com/mrfoe7/ru-doc-code"
)

func TestValidate(t *testing.T) {
	t.Parallel()

	t.Run("invalid ogrn length", func(t *testing.T) {
		testCases := []ru_doc_code.TestCodeCase{
			{
				Code:    "1027700132195",
				Error:   nil,
				IsValid: true,
			},
			{
				Code:    "1027739244741",
				Error:   nil,
				IsValid: true,
			},
			{
				Code:    "102773924",
				Error:   ru_doc_code.ErrInvalidOGRNLength,
				IsValid: false,
			},
			{
				Code:    "10277392447411231",
				Error:   ru_doc_code.ErrInvalidOGRNLength,
				IsValid: false,
			},
		}
		for _, test := range testCases {
			isValid, err := Validate(test.Code)
			assert.Equal(t, test.IsValid, isValid, test.Code)
			assert.Equal(t, true, errors.Is(test.Error, err), test.Code)
		}
	})

	t.Run("invalid ogrn value", func(t *testing.T) {
		testCases := []ru_doc_code.TestCodeCase{
			{
				Code:    "102773??44741",
				Error:   ru_doc_code.ErrInvalidValue,
				IsValid: false,
			},
			{
				Code:    "1027739244742",
				Error:   nil,
				IsValid: false,
			},
			{
				Code:    "10@7739244%42",
				Error:   ru_doc_code.ErrInvalidValue,
				IsValid: false,
			},
			{
				Code:    "1027700132195",
				Error:   nil,
				IsValid: true,
			},
			{
				Code:    "1027739244741",
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
