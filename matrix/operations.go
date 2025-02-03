package matrix

import (
	"github.com/Mehul-Kumar-27/Aayam/concurrency"
	assert "github.com/Mehul-Kumar-27/Aayam/utils"
)

func AddMatrix(mats []Float64Mat, opts ...concurrency.ConcurrencyOptions) (*Float64Mat, error) {
	if len(mats) == 0 {
		return nil, ErrEmptySliceForMatrixAddition
	}
	matsLen := len(mats)
	// Take out the first matrix and use it as the base for validations
	rows, cols := mats[0].Size()
	// Validate whether all the matrixes are of the same dimension
	for i := 1; i < matsLen; i++ {
		r, c := mats[i].Size()
		if !assert.AssertEqual(r, rows) || !assert.AssertEqual(c, cols) {
			return nil, ErrDimensionMismatch
		}
	}
	// Now we are good to go for matrix addition :)
	defaultValue := 0.0
	matOpts := &Float64MatOptions{
		Rows: rows,
		Cols: cols,
		DefaultVal: &defaultValue,
	}
	result := NewMatrix(*matOpts)
	for _, element := range mats {
		matrix := element
		
	}
	return nil, nil
}
