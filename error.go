package zip

import "github.com/pkg/errors"

// error list
var (
	ErrInvalidIndex = errors.New("file index reference is invalid")
	ErrInvalidInput = errors.New("could not create a new reader with nil input")
)
