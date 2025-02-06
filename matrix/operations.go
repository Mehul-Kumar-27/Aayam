package matrix

import (
	"github.com/Mehul-Kumar-27/Aayam/concurrency"
	assert "github.com/Mehul-Kumar-27/Aayam/utils"
)

func AddMatrix(mats []Float64Mat, opts ...*concurrency.ConcurrencyOptions) (*Float64Mat, error) {
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
		Rows:       rows,
		Cols:       cols,
		DefaultVal: &defaultValue,
	}
	result := NewMatrix(*matOpts)
	var concurrencyOpts *concurrency.ConcurrencyOptions
	if len(opts) > 0 && opts[0] != nil {
		concurrencyOpts = opts[0]
	}

	if concurrencyOpts != nil && concurrencyOpts.Enabled {
		var numWorkers int = (result.Rows() * 20) / 100
		if numWorkers < 1 {
			numWorkers = 1
		}
		if concurrencyOpts.Batch_Size != 0 {
			numWorkers = concurrencyOpts.Batch_Size
		}
		err := addMatrixRowParallel(mats, result, numWorkers)
		if err != nil {
			return nil, err
		}
		return result, err
	}
	for _, element := range mats {
		matrix := element
		for r := 0; r < rows; r++ {
			for c := 0; c < cols; c++ {
				result.Data[r][c] += matrix.Data[r][c]
			}
		}
	}
	return result, nil
}

func MultiplyMatrix(leftMatrix Float64Mat, rightMatrix Float64Mat, opts ...*concurrency.ConcurrencyOptions) (*Float64Mat, error) {
	if !assert.AssertEqual(leftMatrix.Columns(), rightMatrix.Rows()) {
		return nil, ErrDimensionMismatch
	}
	leftRows := leftMatrix.Rows()
	leftColumns := leftMatrix.Columns()

	rightColumns := rightMatrix.Columns()

	resultantMatrix := NewMatrix(Float64MatOptions{
		Rows: leftRows,
		Cols: rightColumns,
	})
	// Here we represent the matrix multiplication as the addition of linear combination of columns of matrixes
	for col := 0; col < resultantMatrix.Columns(); col++ {
		// Now we transverse the columns of the left matrix and represent them as a linear combination
		linearCombination := make([]Float64Mat, 0, leftColumns)
		for leftCol := 0; leftCol < leftColumns; leftCol++ {
			leftColData := leftMatrix.GetColumn(leftCol)
			leftColDataAsMatrix := ColumnMatrix(leftColData)
			leftColDataAsMatrix.ScalarMultiplication(rightMatrix.GetVal(leftCol, col))
			linearCombination = append(linearCombination, *leftColDataAsMatrix)
		}
		// Now this linearCombination slice contains the matrix columns and adding then will provide us with this column data
		columnAsMatrix, err := AddMatrix(linearCombination)
		if err != nil {
			return nil, err
		}
		resultantMatrix.SetColumn(columnAsMatrix.GetColumn(0), col)
	}
	return resultantMatrix, nil
}
