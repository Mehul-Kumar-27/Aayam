package vector

import (
	"errors"
)

var (
	ErrEmptyVectorLength = errors.New("length of vectors to add is 0")
	ErrInvalidVectorLength = errors.New("invalid length vector present")
)
