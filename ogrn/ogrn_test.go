package ogrn

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sshaplygin/docs-code/models"
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
				IsValid: true,
			},
			{
				Code:    "1027739244741",
				IsValid: true,
			},
			{
				Code:  "102773924",
				Error: models.ErrInvalidLength,
			},
			{
				Code:  "10277392447411231",
				Error: models.ErrInvalidLength,
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
				Code:  "102773??44741",
				Error: models.ErrInvalidValue,
			},
			{
				Code: "1027739244742",
			},
			{
				Code:  "10@7739244%42",
				Error: models.ErrInvalidValue,
			},
			{
				Code:    "1027700132195",
				IsValid: true,
			},
			{
				Code:    "1027739244741",
				IsValid: true,
			},
		}

		for _, tc := range testCases {
			tc := tc

			isValid, err := Validate(tc.Code)
			assert.Equal(t, tc.IsValid, isValid, tc.Code, tc.IsValid)
			assert.Equal(t, true, errors.Is(tc.Error, err), tc.Code)
		}
	})
}

func Test_Generate(t *testing.T) {
	for i := 0; i < 10; i++ {
		ogrn := Generate()
		isValid, err := Validate(ogrn)
		require.NoError(t, err, fmt.Sprintf("invalid ogrn value: %s", ogrn))

		assert.True(t, isValid)
	}
}

func BenchmarkValidateCorrect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Validate("1027700132195")
	}
}
func BenchmarkGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Generate()
	}
}
