package okato

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sshaplygin/docs-code/models"
)

func Test_Validete(t *testing.T) {
	type testCase struct {
		Code    string
		IsValid bool
		Error   error
	}

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
			Code:    "17 205 000 000",
			IsValid: true,
		},
		{
			Code:    "01 201 802 003",
			IsValid: true,
		},
		{
			Code:    "45286560000",
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
}

func Test_Generate(t *testing.T) {
	require.Panics(t, func() {
		Generate()
	})
}
