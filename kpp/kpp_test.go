package kpp

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	ru_doc_code "github.com/mrfoe7/go-codes-validator"
)

func TestValidate(t *testing.T) {
	t.Parallel()

	t.Run("invalid kpp length", func(t *testing.T) {
		testCases := []ru_doc_code.TestCodeCase{
			{
				Code:    "1234567888776",
				Error:   ru_doc_code.ErrInvalidKPPLength,
				IsValid: false,
			},
			{
				Code:    "044525",
				Error:   ru_doc_code.ErrInvalidKPPLength,
				IsValid: false,
			},
			{
				Code:    "773643301",
				Error:   nil,
				IsValid: true,
			},
			{
				Code:    "773643001",
				Error:   nil,
				IsValid: true,
			},
		}
		for _, test := range testCases {
			isValid, err := Validate(test.Code)
			assert.Equal(t, isValid, test.IsValid, test.Code)
			assert.Equal(t, true, errors.Is(test.Error, err), test.Code)
		}
	})

	t.Run("invalid kpp value", func(t *testing.T) {
		testCases := []ru_doc_code.TestCodeCase{
			{
				Code:    "773$N3301",
				Error:   ru_doc_code.ErrInvalidValue,
				IsValid: false,
			},
			{
				Code:    "7736#3&01",
				Error:   ru_doc_code.ErrInvalidValue,
				IsValid: false,
			},
			{
				Code:    "773643301",
				Error:   nil,
				IsValid: true,
			},
			{
				Code:    "773643001",
				Error:   nil,
				IsValid: true,
			},
		}
		for _, test := range testCases {
			isValid, err := Validate(test.Code)
			assert.Equal(t, isValid, test.IsValid, test.Code)
			assert.Equal(t, true, errors.Is(test.Error, err), test.Code)
		}
	})
}
