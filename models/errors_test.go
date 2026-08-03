package models

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCommonError_Error(t *testing.T) {
	err := &CommonError{Method: "inn", Err: ErrInvalidLength}
	assert.Equal(t, "inn: invalid length", err.Error())
}

func TestCommonError_Unwrap(t *testing.T) {
	base := ErrInvalidValue
	err := &CommonError{Method: "bik", Err: base}

	require.Equal(t, base, err.Unwrap())
}

func TestCommonError_ErrorsIs(t *testing.T) {
	t.Run("direct", func(t *testing.T) {
		err := &CommonError{Method: "kpp", Err: ErrInvalidLength}
		assert.True(t, errors.Is(err, ErrInvalidLength))
		assert.False(t, errors.Is(err, ErrInvalidValue))
	})

	t.Run("wrapped", func(t *testing.T) {
		err := fmt.Errorf("parse model: %w", &CommonError{Method: "snils", Err: ErrInvalidValue})
		assert.True(t, errors.Is(err, ErrInvalidValue))
		assert.False(t, errors.Is(err, ErrInvalidLength))
	})
}
