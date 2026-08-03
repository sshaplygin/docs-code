package docs_code

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Generate(t *testing.T) {
	inn, err := Generate(INN)
	require.NoError(t, err)

	isValid, err := Validate(INN, inn)
	require.NoError(t, err)

	require.True(t, isValid)
}

func Test_Generate_Unsupported(t *testing.T) {
	_, err := Generate(DocType(100500))
	require.ErrorIs(t, err, ErrUnsupportedDocType)
}
