package bik

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sshaplygin/docs-code/models"
)

func TestValidate(t *testing.T) {
	t.Parallel()

	type testCase struct {
		Code    string
		IsValid bool
		Error   error
	}

	t.Run("invalid bik length", func(t *testing.T) {
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

		for i, tc := range testCases {
			tc := tc

			isValid, err := Validate(tc.Code)
			assert.Equal(t, tc.IsValid, isValid, tc.Code)
			if err != nil {
				assert.ErrorAs(t, err, &tc.Error, fmt.Sprintf("invalid test case %d: input: %s", i, tc.Code))
			} else {
				assert.NoError(t, err, fmt.Sprintf("invalid test case %d: input: %s", i, tc.Code))
			}
		}
	})

	t.Run("invalid bik value", func(t *testing.T) {
		type testCase struct {
			Code    string
			IsValid bool
			Error   error
		}

		testCases := []testCase{
			{
				Code:    "0445?5226",
				Error:   models.ErrInvalidValue,
				IsValid: false,
			},
			{
				Code:    "054525225",
				Error:   ErrInvalidCountryCode,
				IsValid: false,
			},
			{
				Code:    "104525225",
				Error:   ErrInvalidCountryCode,
				IsValid: false,
			},
			{
				Code:    "044#55#25",
				Error:   models.ErrInvalidValue,
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
		for i, tc := range testCases {
			tc := tc

			isValid, err := Validate(tc.Code)
			assert.Equal(t, tc.IsValid, isValid, tc.Code, tc.IsValid)
			if err != nil {
				assert.ErrorAs(t, err, &tc.Error, fmt.Sprintf("invalid test case %d: input: %s", i, tc.Code))
			} else {
				assert.Empty(t, err, fmt.Sprintf("invalid test case %d: input: %s", i, tc.Code))
			}
		}
	})
}

func Test_Generate(t *testing.T) {
	require.Panics(t, func() {
		Generate()
	})
}
