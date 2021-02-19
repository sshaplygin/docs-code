package inn

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	ru_doc_code "github.com/mrfoe7/ru-doc-code"
)

func TestValidate(t *testing.T) {
	t.Parallel()

	t.Run("invalid inn length", func(t *testing.T) {
		testCases := []ru_doc_code.TestCodeCase{
			{
				Code:    "12345678",
				Error:   ru_doc_code.ErrInvalidINNLength,
				IsValid: false,
			},
			{
				Code:    "9876543211123",
				Error:   ru_doc_code.ErrInvalidINNLength,
				IsValid: false,
			},
			{
				Code:    "7707083893",
				Error:   nil,
				IsValid: true,
			},
			{
				Code:    "526317984689",
				Error:   nil,
				IsValid: true,
			},
		}
		for _, test := range testCases {
			isValid, err := Validate(test.Code)
			assert.Equal(t, test.IsValid, isValid, test.Code)
			assert.Equal(t, true, errors.Is(test.Error, err), test.Code)
		}
	})

	t.Run("invalid inn value", func(t *testing.T) {
		testCases := []ru_doc_code.TestCodeCase{
			{
				Code:    "77$7083893",
				Error:   ru_doc_code.ErrInvalidValue,
				IsValid: false,
			},
			{
				Code:    "98754321N123",
				Error:   ru_doc_code.ErrInvalidValue,
				IsValid: false,
			},
			{
				Code:    "9854132d1123",
				Error:   ru_doc_code.ErrInvalidValue,
				IsValid: false,
			},
			{
				Code:    "7707083893",
				Error:   nil,
				IsValid: true,
			},
			{
				Code:    "526317984689",
				Error:   nil,
				IsValid: true,
			},
		}
		for _, test := range testCases {
			isValid, err := Validate(test.Code)
			assert.Equal(t, test.IsValid, isValid, test.Code)
			assert.Equal(t, true, errors.Is(test.Error, err), test.Code)
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

			require.True(t, isValid)
			require.NoError(t, err)
		}
	})

	t.Run("generate physical inn", func(t *testing.T) {
		var inn string

		for i := 0; i < 10; i++ {
			inn = GeneratePhysical()
			isValid, err := Validate(inn)

			require.True(t, isValid)
			require.NoError(t, err)
		}
	})

	t.Run("generate random inn", func(t *testing.T) {
		var inn string

		for i := 0; i < 10; i++ {
			inn = Generate()
			isValid, err := Validate(inn)

			require.True(t, isValid)
			require.NoError(t, err)
		}
	})

	t.Run("generate random digits ", func(t *testing.T) {
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

		for _, test := range tests {
			digits = ru_doc_code.RandomDigits(test.len)

			assert.True(t, digits >= test.min && digits <= test.max)
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
