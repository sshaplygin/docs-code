package inn

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sshaplygin/docs-code/models"
	"github.com/sshaplygin/docs-code/utils"
)

func TestValidate(t *testing.T) {
	t.Parallel()

	t.Run("invalid inn length", func(t *testing.T) {
		type testCase struct {
			Code    string
			IsValid bool
			Error   error
		}

		testCases := []testCase{
			{
				Code:    "12345678",
				Error:   models.ErrInvalidLength,
				IsValid: false,
			},
			{
				Code:    "9876543211123",
				Error:   models.ErrInvalidLength,
				IsValid: false,
			},
			{
				Code:    "7707083893",
				IsValid: true,
			},
			{
				Code:    "526317984689",
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
				assert.Empty(t, err, fmt.Sprintf("invalid test case %d: input: %s", i, tc.Code))
			}
		}
	})

	t.Run("invalid inn value", func(t *testing.T) {
		type testCase struct {
			Code    string
			IsValid bool
			Error   error
		}

		testCases := []testCase{
			{
				Code:    "77$7083893",
				Error:   models.ErrInvalidValue,
				IsValid: false,
			},
			{
				Code:  "98754321N123",
				Error: models.ErrInvalidValue,
			},
			{
				Code:  "9854132d1123",
				Error: models.ErrInvalidValue,
			},
			{
				Code:    "7707083893",
				IsValid: true,
			},
			{
				Code:    "526317984689",
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
				assert.Empty(t, err, fmt.Sprintf("invalid test case %d: input: %s", i, tc.Code))
			}
		}
	})
}

func TestGenerate(t *testing.T) {
	t.Parallel()

	t.Run("generate legal inn", func(t *testing.T) {
		var inn string

		for i := 0; i < 10; i++ {
			inn = GenerateLegal()
			isValid, err := Validate(inn)
			require.NoError(t, err)

			require.True(t, isValid)
		}
	})

	t.Run("generate physical inn", func(t *testing.T) {
		var inn string

		for i := 0; i < 10; i++ {
			inn = GeneratePhysical()
			isValid, err := Validate(inn)
			require.NoError(t, err)

			require.True(t, isValid)
		}
	})

	t.Run("generate random inn", func(t *testing.T) {
		var inn string

		for i := 0; i < 10; i++ {
			inn = Generate()
			isValid, err := Validate(inn)
			require.NoError(t, err)

			require.True(t, isValid)
		}
	})

	t.Run("generate random digits", func(t *testing.T) {
		tests := []struct {
			len int
			min int64
			max int64
		}{
			{
				-5,
				0,
				9,
			},
			{
				-10,
				0,
				9,
			},
			{
				1,
				0,
				9,
			},
			{
				3,
				100,
				999,
			},
		}
		var digits int64

		for _, tc := range tests {
			tc := tc

			digits = utils.RandomDigits(tc.len)
			assert.True(t, digits >= tc.min && digits <= tc.max)
		}
	})
}

func BenchmarkGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Generate()
	}
}

func BenchmarkGenerateLegal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateLegal()
	}
}

func BenchmarkGeneratePhysical(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GeneratePhysical()
	}
}
