package docs_code

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Generate(t *testing.T) {
	inn := Generate(INN)
	isValid, err := Validate(INN, inn)
	require.NoError(t, err)

	require.True(t, isValid)
}

func Test_Generate_Unsupported(t *testing.T) {
	require.Panics(t, func() {
		Generate(DocType(100500))
	})
}
