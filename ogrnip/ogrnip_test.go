package ogrnip

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sshaplygin/docs-code/models"
)

func TestValidate(t *testing.T) {
	t.Parallel()

	t.Run("invalid ogrnip length", func(t *testing.T) {
		type testCase struct {
			Code    string
			IsValid bool
			Error   error
		}

		testCases := []testCase{
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
				Error:   models.ErrInvalidLength,
				IsValid: false,
			},
			{
				Code:    "3045001111236000157",
				Error:   models.ErrInvalidLength,
				IsValid: false,
			},
		}
		for i, tc := range testCases {
			tc := tc

			isValid, err := Validate(tc.Code)
			assert.Equal(t, tc.IsValid, isValid, tc.Code)
			if err != nil {
				assert.ErrorAs(t, err, &tc.Error, fmt.Sprintf("invalid test case %d: input: %s", i, tc.Code))
			} else {
				assert.Empty(t, err, fmt.Sprintf("invalid test case %d: input: %s", i, tc.Code))
			}
		}
	})

	t.Run("invalid ogrnip value", func(t *testing.T) {
		type testCase struct {
			Code    string
			IsValid bool
			Error   error
		}

		testCases := []testCase{
			{
				Code:    "312502??4600034",
				Error:   models.ErrInvalidValue,
				IsValid: false,
			},
			{
				Code:    "304500116000158",
				Error:   nil,
				IsValid: false,
			},
			{
				Code:    "512502904600034",
				Error:   models.ErrInvalidValue,
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
