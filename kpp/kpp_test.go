package kpp

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sshaplygin/ru-doc-code/models"
)

//nolint:dupl
func TestValidate(t *testing.T) {
	t.Parallel()

	t.Run("invalid kpp length", func(t *testing.T) {
		type testCase struct {
			Code    string
			IsValid bool
			Error   error
		}

		testCases := []testCase{
			{
				Code:    "1234567888776",
				Error:   models.ErrInvalidLength,
				IsValid: false,
			},
			{
				Code:    "044525",
				Error:   models.ErrInvalidLength,
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
		for i, tc := range testCases {
			tc := tc

			isValid, err := Validate(tc.Code)
			assert.Equal(t, isValid, tc.IsValid, fmt.Sprintf("invalid test case %d: input: %s", i, tc.Code))
			if err != nil {
				assert.ErrorAs(t, err, &tc.Error, fmt.Sprintf("invalid test case %d: input: %s", i, tc.Code))
			} else {
				assert.Empty(t, err, fmt.Sprintf("invalid test case %d: input: %s", i, tc.Code))
			}
		}
	})

	t.Run("invalid kpp value", func(t *testing.T) {
		type testCase struct {
			Code    string
			IsValid bool
			Error   error
		}

		testCases := []testCase{
			{
				Code:    "773$N3301",
				Error:   models.ErrInvalidValue,
				IsValid: false,
			},
			{
				Code:    "7736#3&01",
				Error:   models.ErrInvalidValue,
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
		for i, tc := range testCases {
			tc := tc

			isValid, err := Validate(tc.Code)
			assert.Equal(t, isValid, tc.IsValid, fmt.Sprintf("invalid test case %d: input: %s", i, tc.Code))
			if err != nil {
				assert.ErrorAs(t, err, &tc.Error, fmt.Sprintf("invalid test case %d: input: %s", i, tc.Code))
			} else {
				assert.Empty(t, err, fmt.Sprintf("invalid test case %d: input: %s", i, tc.Code))
			}
		}
	})

	t.Run("invalid registration reason code", func(t *testing.T) {
		type testCase struct {
			Code    string
			IsValid bool
			Error   error
		}

		testCases := []testCase{
			{
				Code:    "773643301",
				Error:   ErrRegistrationReasonCode,
				IsValid: false,
			},
			{
				Code:    "773642301",
				Error:   nil,
				IsValid: true,
			},
		}
		for i, tc := range testCases {
			tc := tc

			isValid, err := Validate(tc.Code)
			assert.Equal(t, isValid, tc.IsValid, fmt.Sprintf("invalid test case %d: input: %s", i, tc.Code))
			if err != nil {
				assert.ErrorAs(t, err, &tc.Error, fmt.Sprintf("invalid test case %d: input: %s", i, tc.Code))
			} else {
				assert.Empty(t, err, fmt.Sprintf("invalid test case %d: input: %s", i, tc.Code))
			}
		}
	})
}
