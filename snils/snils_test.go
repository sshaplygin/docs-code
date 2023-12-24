package snils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sshaplygin/docs-code/models"
)

func TestValidate(t *testing.T) {
	t.Parallel()

	t.Run("invalid snils length", func(t *testing.T) {
		type testCase struct {
			Code    string
			IsValid bool
			Error   error
		}

		testCases := []testCase{
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
				Error:   models.ErrInvalidLength,
				IsValid: false,
			},
			{
				Code:    "112-233 95",
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

	t.Run("invalid snils value", func(t *testing.T) {
		type testCase struct {
			Code    string
			IsValid bool
			Error   error
		}

		testCases := []testCase{
			{
				Code:    "112-233?445 95",
				Error:   ErrInvalidFormattedLength,
				IsValid: false,
			},
			{
				Code:    "1M2-234-445 95",
				Error:   models.ErrInvalidValue,
				IsValid: false,
			},
			{
				Code:    "112-233-445 98",
				Error:   nil,
				IsValid: false,
			},
			{
				Code:    "112-233-445#95",
				Error:   ErrInvalidFormattedLength,
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
