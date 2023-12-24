package oktmo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Validete(t *testing.T) {
	require.Panics(t, func() {
		Validate("")
	})
}

func Test_Generate(t *testing.T) {
	require.Panics(t, func() {
		Generate()
	})
}
