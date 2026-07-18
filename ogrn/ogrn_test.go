package ogrn

import (
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
			isValid, err := Validate(tc.Code)
			assert.Equal(t, tc.IsValid, isValid, tc.Code)
			if err != nil {
				assert.ErrorIs(t, err, tc.Error, fmt.Sprintf("invalid test case %d: input: %s", i, tc.Code))
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

		for i, tc := range testCases {
			isValid, err := Validate(tc.Code)
			assert.Equal(t, tc.IsValid, isValid, tc.Code)
			if err != nil {
				assert.ErrorIs(t, err, tc.Error, fmt.Sprintf("invalid test case %d: input: %s", i, tc.Code))
			} else {
				assert.Empty(t, err, fmt.Sprintf("invalid test case %d: input: %s", i, tc.Code))
			}
		}
	})

	t.Run("ogrn type detection", func(t *testing.T) {
		type testCase struct {
			Code    string
			IsValid bool
			Error   error
		}

		testCases := []testCase{
			{
				// government OGRN (leading 2) is now accepted
				Code:    "2027700132194",
				IsValid: true,
			},
			{
				// government OGRN (leading 9) is now accepted
				Code:    "9027700132198",
				IsValid: true,
			},
			{
				// leading 3 is an individual-entrepreneur (physical) code type
				// that belongs to the ogrnip package, so it must be rejected here
				Code:  "3027700132195",
				Error: ErrInvalidCodeType,
			},
			{
				Code:  "",
				Error: models.ErrInvalidLength,
			},
			{
				Code:  "A027700132195",
				Error: models.ErrInvalidValue,
			},
		}

		for i, tc := range testCases {
			isValid, err := Validate(tc.Code)
			assert.Equal(t, tc.IsValid, isValid, tc.Code)
			if tc.Error != nil {
				assert.ErrorIs(t, err, tc.Error, fmt.Sprintf("invalid test case %d: input: %s", i, tc.Code))
			} else {
				assert.NoError(t, err, fmt.Sprintf("invalid test case %d: input: %s", i, tc.Code))
			}
		}
	})
}

func Test_Generate(t *testing.T) {
	for range 10 {
		ogrn := Generate()
		isValid, err := Validate(ogrn)
		require.NoError(t, err, fmt.Sprintf("invalid ogrn value: %s", ogrn))

		assert.True(t, isValid)
	}
}

func BenchmarkValidateCorrect(b *testing.B) {
	for b.Loop() {
		_, _ = Validate("1027700132195")
	}
}
func BenchmarkGenerate(b *testing.B) {
	for b.Loop() {
		Generate()
	}
}
