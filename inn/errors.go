package inn

import "errors"

var (
	// ErrNilINN try call methods for nil inn struct
	ErrNilINN = errors.New("nil inn struct")

	// ErrInvalidCheckSumsValue invalid check sums rune value
	ErrInvalidCheckSumsValue = errors.New("invalid check sums rune value")
)
