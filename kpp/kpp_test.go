package kpp

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	ru_doc_code "github.com/sshaplygin/ru-doc-code"
)

func TestValidate(t *testing.T) {
	t.Parallel()

	t.Run("invalid kpp length", func(t *testing.T) {
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
				Code:    "773642301",
				Error:   nil,
				IsValid: true,
			},
			{
				Code:    "773642001",
				Error:   nil,
				IsValid: true,
			},
		}
		for i, test := range testCases {
			isValid, err := Validate(test.Code)
			assert.Equal(t, isValid, test.IsValid, fmt.Sprintf("invalid test case %d: input: %s", i, test.Code))
			if err != nil {
				assert.True(t, errors.As(err, &test.Error), fmt.Sprintf("invalid test case %d: input: %s", i, test.Code))
			} else {
				assert.Empty(t, err, fmt.Sprintf("invalid test case %d: input: %s", i, test.Code))
			}
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
				Code:    "773642301",
				Error:   nil,
				IsValid: true,
			},
			{
				Code:    "773642001",
				Error:   nil,
				IsValid: true,
			},
		}
		for i, test := range testCases {
			isValid, err := Validate(test.Code)
			assert.Equal(t, isValid, test.IsValid, fmt.Sprintf("invalid test case %d: input: %s", i, test.Code))
			if err != nil {
				assert.True(t, errors.As(err, &test.Error), fmt.Sprintf("invalid test case %d: input: %s", i, test.Code))
			} else {
				assert.Empty(t, err, fmt.Sprintf("invalid test case %d: input: %s", i, test.Code))
			}
		}
	})

	t.Run("invalid registration reason code", func(t *testing.T) {
		testCases := []ru_doc_code.TestCodeCase{
			{
				Code:    "773643301",
				Error:   ru_doc_code.ErrRegistrationReasonCode,
				IsValid: false,
			},
			{
				Code:    "773642301",
				Error:   nil,
				IsValid: true,
			},
		}
		for i, test := range testCases {
			isValid, err := Validate(test.Code)
			assert.Equal(t, isValid, test.IsValid, fmt.Sprintf("invalid test case %d: input: %s", i, test.Code))
			if err != nil {
				assert.True(t, errors.As(err, &test.Error), fmt.Sprintf("invalid test case %d: input: %s", i, test.Code))
			} else {
				assert.Empty(t, err, fmt.Sprintf("invalid test case %d: input: %s", i, test.Code))
			}
		}
	})
}
