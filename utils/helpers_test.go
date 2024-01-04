package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func Test_StrCode_ValidInput(t *testing.T) {
	type testCase struct {
		name   string
		val    int
		length int
		want   string
	}

	tests := []testCase{
		{
			"nil",
			1,
			2,
			"01",
		},
		{
			"nil",
			1,
			3,
			"001",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, StrCode(tc.val, tc.length))
		})
	}
}

func Test_Generate_InvalidInput(t *testing.T) {
	type testCase struct {
		name   string
		val    int
		length int
	}

	tests := []testCase{
		{
			"nil",
			0,
			0,
		},
		{
			"nil",
			100,
			2,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			require.Panics(t, func() {
				StrCode(tc.val, tc.length)
			})
		})
	}
}
