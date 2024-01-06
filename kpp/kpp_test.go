package kpp

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sshaplygin/docs-code/models"
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
				Code:    "775001001",
				IsValid: true,
			},
			{
				Code:    "773642301",
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

func Test_Generate(t *testing.T) {
	kpp := Generate()
	isValid, err := Validate(kpp)
	require.NoError(t, err, fmt.Sprintf("invalid kpp value: %s", kpp))

	assert.True(t, isValid)
}
