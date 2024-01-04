package okato

import "errors"

var (
	// ErrNilOKATO try call methods for nil okato struct
	ErrNilOKATO = errors.New("nil okato struct")

	// ErrInvalidCode invalid okato code
	ErrInvalidCode = errors.New("invalid okato code")

	// ErrFirstLevelCode
	ErrFirstLevelCode = errors.New("")

	// ErrSecondLevelCode
	ErrSecondLevelCode = errors.New("")

	// ErrThirdLevelCode
	ErrThirdLevelCode = errors.New("")

	// ErrFourthLevelCode
	ErrFourthLevelCode = errors.New("")
)
