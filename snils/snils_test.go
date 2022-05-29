package snils

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	ru_doc_code "github.com/sshaplygin/ru-doc-code"
)

func TestValidate(t *testing.T) {
	t.Parallel()

	t.Run("invalid snils length", func(t *testing.T) {
		testCases := []ru_doc_code.TestCodeCase{
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
				Error:   ru_doc_code.ErrInvalidLength,
				IsValid: false,
			},
			{
				Code:    "112-233 95",
				Error:   ru_doc_code.ErrInvalidLength,
				IsValid: false,
			},
		}
		for i, test := range testCases {
			isValid, err := Validate(test.Code)
			assert.Equal(t, test.IsValid, isValid, test.Code)
			if err != nil {
				assert.True(t, errors.As(err, &test.Error), fmt.Sprintf("invalid test case %d: input: %s", i, test.Code))
			} else {
				assert.Empty(t, err, fmt.Sprintf("invalid test case %d: input: %s", i, test.Code))
			}
		}
	})

	t.Run("invalid snils value", func(t *testing.T) {
		testCases := []ru_doc_code.TestCodeCase{
			{
				Code:    "112-233?445 95",
				Error:   ru_doc_code.ErrInvalidFormattedSNILSLength,
				IsValid: false,
			},
			{
				Code:    "1M2-234-445 95",
				Error:   ru_doc_code.ErrInvalidValue,
				IsValid: false,
			},
			{
				Code:    "112-233-445 98",
				Error:   nil,
				IsValid: false,
			},
			{
				Code:    "112-233-445#95",
				Error:   ru_doc_code.ErrInvalidFormattedSNILSLength,
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
		for i, test := range testCases {
			isValid, err := Validate(test.Code)
			assert.Equal(t, test.IsValid, isValid, test.Code, test.IsValid)
			if err != nil {
				assert.True(t, errors.As(err, &test.Error), fmt.Sprintf("invalid test case %d: input: %s", i, test.Code))
			} else {
				assert.Empty(t, err, fmt.Sprintf("invalid test case %d: input: %s", i, test.Code))
			}
		}
	})
}
