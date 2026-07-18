package okato

import "errors"

var (
	// ErrNilOKATO try call methods for nil okato struct
	ErrNilOKATO = errors.New("nil okato struct")

	// ErrInvalidCode invalid okato code
	ErrInvalidCode = errors.New("invalid okato code")

	// ErrFirstLevelCode invalid first level (subject of the Russian Federation) code
	ErrFirstLevelCode = errors.New("invalid okato first level code")

	// ErrSecondLevelCode invalid second level code
	ErrSecondLevelCode = errors.New("invalid okato second level code")

	// ErrThirdLevelCode invalid third level code
	ErrThirdLevelCode = errors.New("invalid okato third level code")

	// ErrFourthLevelCode invalid fourth level code
	ErrFourthLevelCode = errors.New("invalid okato fourth level code")

	// ErrNotImplemented feature is not implemented yet
	ErrNotImplemented = errors.New("not implemented")
)
