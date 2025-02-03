package matrix

import (
	"errors"
)

var (
	ErrDimensionMismatch           = errors.New("matrix dimensions do not match for multiplication")
	ErrEmptySliceForMatrixAddition = errors.New("matrix slice is empty")
)
