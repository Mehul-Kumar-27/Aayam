package vector

import (
	"sync"

	"github.com/Mehul-Kumar-27/Aayam/concurrency"
)

// AddFloat64Vectors adds multiple IntegerVectors together and returns a pointer to the resultant vector and an error if any.
// It takes a variadic number of IntegerVector arguments and performs the following steps:
// 1. Checks if the input vectors are non-empty. If empty, returns ErrEmptyVectorLength.
// 2. Ensures all vectors are of the same size. If not, returns ErrInvalidVectorLength.
// 3. Iterates over each vector and adds corresponding elements to the initial vector.
// The function returns a pointer to the resultant IntegerVector and nil if successful, or an error if any of the checks fail.
func AddFloat64Vectors(vectors []Float64Vec, opts ...*concurrency.ConcurrencyOptions) (*Float64Vec, error) {

	// Get the size of the first vector
	size := vectors[0].Size()
	if size == 0 {
		return nil, ErrEmptyVectorLength
	}
	var vectors_len int = len(vectors)
	// Check whether all the vectors are of the same size
	for i := 1; i < vectors_len; i++ {
		if vectors[i].Size() != size {
			return nil, ErrInvalidVectorLength
		}
	}

	// Create a new vector with the same size and default value of 0
	resultant_vector := NewVector(Float64VecOptions{Size: size})

	///Opts provided means we would use the concurrency support
	var concurrencyOpts *concurrency.ConcurrencyOptions
	if len(opts) > 0 && opts[0] != nil {
		concurrencyOpts = opts[0]
	}
	if concurrencyOpts != nil && concurrencyOpts.Enabled {
		var batchSize int = (len(vectors) * 20) / 100
		if batchSize < 1 {
			batchSize = 1
		}
		if concurrencyOpts.Batch_Size != 0 {
			batchSize = concurrencyOpts.Batch_Size
		}
		err := addVectorsConcurrently(vectors, resultant_vector, batchSize)
		if err != nil {
			return nil, err
		}
		return resultant_vector, nil
	}
	// Add all vectors element-wise
	for _, vec := range vectors {
		for index := 0; index < size; index++ {
			resultant_vector.Data[index] += vec.Data[index]
		}
	}

	return resultant_vector, nil
}

func addVectorsConcurrently(vectors []Float64Vec, resultant_vector *Float64Vec, batchSize int) error {
	total_vectors := len(vectors)
	numBatches := (total_vectors + batchSize - 1) / batchSize
	size := resultant_vector.Size()
	var wg sync.WaitGroup

	resultChan := make(chan *Float64Vec, numBatches)

	for i := 0; i < numBatches; i++ {
		wg.Add(1)
		go func(batchIndex int) {
			defer wg.Done()
			start := batchIndex * batchSize
			end := start + batchSize
			if end > total_vectors {
				end = total_vectors
			}

			// Initialize the vector to store this batch result
			batchResult := NewVector(Float64VecOptions{
				Size: resultant_vector.Size(),
			})

			for index := start; index < end; index++ {
				vec := vectors[index]
				for i := 0; i < size; i++ {
					batchResult.Data[i] += vec.GetVal(i)
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
		for i, val := range batchResult.Data {
			resultant_vector.Data[i] += val
		}
	}

	return nil
}
