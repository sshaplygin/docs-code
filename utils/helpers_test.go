package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SliceToInt(t *testing.T) {
	type testCase struct {
		name string
		in   []int
		want int
	}

	tests := []testCase{
		{
			"nil",
			nil,
			0,
		},
		{
			"102",
			[]int{1, 0, 2},
			102,
		},
		{
			"with zero prefix 02",
			[]int{0, 2},
			2,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, SliceToInt(tc.in))
		})
	}
}
