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
				Code:  "1234567888776",
				Error: models.ErrInvalidLength,
			},
			{
				Code:  "044525",
				Error: models.ErrInvalidLength,
			},
			{
				Code:    "044525225",
				IsValid: true,
			},
			{
				Code:    "044525012",
				IsValid: true,
			},
		}

		for i, tc := range testCases {
			tc := tc

			isValid, err := Validate(tc.Code)
			if err != nil {
				require.ErrorAs(t, err, &tc.Error, fmt.Sprintf("invalid test case %d: input: %s", i, tc.Code))
			} else {
				require.NoError(t, err, fmt.Sprintf("invalid test case %d: input: %s", i, tc.Code))
			}

			assert.Equal(t, tc.IsValid, isValid, tc.Code)
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
				Code:  "0445?5226",
				Error: models.ErrInvalidValue,
			},
			{
				Code:  "054525225",
				Error: ErrInvalidCountryCode,
			},
			{
				Code:  "104525225",
				Error: ErrInvalidCountryCode,
			},
			{
				Code:  "044#55#25",
				Error: models.ErrInvalidValue,
			},
			{
				Code:    "044525225",
				IsValid: true,
			},
			{
				Code:    "044525012",
				IsValid: true,
			},
		}
		for i, tc := range testCases {
			tc := tc

			isValid, err := Validate(tc.Code)
			if err != nil {
				require.ErrorAs(t, err, &tc.Error, fmt.Sprintf("invalid test case %d: input: %s", i, tc.Code))
			} else {
				require.Empty(t, err, fmt.Sprintf("invalid test case %d: input: %s", i, tc.Code))
			}

			assert.Equal(t, tc.IsValid, isValid, tc.Code, tc.IsValid)
		}
	})
}

func Test_Generate(t *testing.T) {
	bik := Generate()
	isValid, err := Validate(bik)

	require.NoError(t, err, fmt.Sprintf("invalid bik value: %s", bik))
	require.True(t, isValid)
}

func Test_Exists(t *testing.T) {
	is, err := Exists("044525677") // АО "Яндекс Банк".
	require.NoError(t, err)

	assert.True(t, is)
}
