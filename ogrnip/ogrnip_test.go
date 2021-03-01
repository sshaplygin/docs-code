package ogrnip

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	ru_doc_code "github.com/mrfoe7/ru-doc-code"
)

func TestValidate(t *testing.T) {
	t.Parallel()

	t.Run("invalid ogrnip length", func(t *testing.T) {
		testCases := []ru_doc_code.TestCodeCase{
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
				Error:   ru_doc_code.ErrInvalidLength,
				IsValid: false,
			},
			{
				Code:    "3045001111236000157",
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

	t.Run("invalid ogrnip value", func(t *testing.T) {
		testCases := []ru_doc_code.TestCodeCase{
			{
				Code:    "312502??4600034",
				Error:   ru_doc_code.ErrInvalidValue,
				IsValid: false,
			},
			{
				Code:    "304500116000158",
				Error:   nil,
				IsValid: false,
			},
			{
				Code:    "512502904600034",
				Error:   ru_doc_code.ErrInvalidValue,
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
