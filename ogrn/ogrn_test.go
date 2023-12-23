package ogrn

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sshaplygin/ru-doc-code/models"
)

func TestValidate(t *testing.T) {
	t.Parallel()

	t.Run("invalid ogrn length", func(t *testing.T) {
		type testCase struct {
			Code    string
			IsValid bool
			Error   error
		}

		testCases := []testCase{
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
				Error:   models.ErrInvalidLength,
				IsValid: false,
			},
			{
				Code:    "10277392447411231",
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

	t.Run("invalid ogrn value", func(t *testing.T) {
		type testCase struct {
			Code    string
			IsValid bool
			Error   error
		}

		testCases := []testCase{
			{
				Code:    "102773??44741",
				Error:   models.ErrInvalidValue,
				IsValid: false,
			},
			{
				Code:    "1027739244742",
				Error:   nil,
				IsValid: false,
			},
			{
				Code:    "10@7739244%42",
				Error:   models.ErrInvalidValue,
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
