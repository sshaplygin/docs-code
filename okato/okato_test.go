package okato

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Validete(t *testing.T) {
	require.Panics(t, func() {
		_, err := Validate("")
		require.NoError(t, err)
	})
}

func Test_Generate(t *testing.T) {
	require.Panics(t, func() {
		Generate()
	})
}
