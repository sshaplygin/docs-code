package docs_code

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Validate(t *testing.T) {
	isValid, err := Validate(INN, "7707083893")
	require.NoError(t, err)

	require.True(t, isValid)
}

func Test_Validate_Unsupported(t *testing.T) {
	require.Panics(t, func() {
		Validate(DocType(100500), "100500")
	})
}
