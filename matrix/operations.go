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
		return nil, nil
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

func addMatrixConcurrently(mats []Float64Mat, resultantMat *Float64Mat, numWorkers int) error {
	totalMats := len(mats)
	numBatches := (totalMats + numWorkers - 1) / numWorkers
	var wg sync.WaitGroup

	resultChan := make(chan *Float64Mat, numBatches)
	for i := 0; i < numBatches; i++ {
		wg.Add(1)
		go func(batchIndex int) {
			defer wg.Done()
			start := batchIndex * numWorkers
			end := start + numWorkers
			if end > totalMats {
				end = totalMats
			}
			batchResult := NewMatrix(Float64MatOptions{
				Rows: resultantMat.Rows(),
				Cols: resultantMat.Columns(),
			})

			for index := start; index < end; index++ {
				matrix := mats[index]
				for r := 0; r < batchResult.Rows(); r++ {
					for c := 0; c < batchResult.Columns(); c++ {
						batchResult.Data[r][c] += matrix.Data[r][c]
					}
				}
			}
			resultChan <- batchResult
		}(i)
	}
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for batchResult := range resultChan {
		for r := 0; r < resultantMat.Rows(); r++ {
			for c := 0; c < resultantMat.Columns(); c++ {
				resultantMat.Data[r][c] += batchResult.Data[r][c]
			}
		}
	}
	return nil
}
