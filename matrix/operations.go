package matrix

import (
	"sync"

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

func addMatrixRowParallel(mats []Float64Mat, resultantMat *Float64Mat, numWorkers int) error {
	rows := resultantMat.Rows()
	cols := resultantMat.Columns()
	var wg sync.WaitGroup

	// Create a channel to distribute row indices to workers.
	rowChan := make(chan int, rows)

	// Start worker goroutines.
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Each worker processes rows sent through the channel.
			for r := range rowChan {
				// For each column in the row, compute the sum over all matrices.
				for c := 0; c < cols; c++ {
					var sum float64
					for _, mat := range mats {
						sum += mat.Data[r][c]
					}
					resultantMat.Data[r][c] = sum
				}
			}
		}()
	}

	for r := 0; r < rows; r++ {
		rowChan <- r
	}
	close(rowChan)

	wg.Wait()
	return nil
}
