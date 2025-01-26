package vector

import "github.com/Mehul-Kumar-27/Aayam/concurrency"

// AddFloat64Vectors adds multiple IntegerVectors together and returns a pointer to the resultant vector and an error if any.
// It takes a variadic number of IntegerVector arguments and performs the following steps:
// 1. Checks if the input vectors are non-empty. If empty, returns ErrEmptyVectorLength.
// 2. Ensures all vectors are of the same size. If not, returns ErrInvalidVectorLength.
// 3. Iterates over each vector and adds corresponding elements to the initial vector.
// The function returns a pointer to the resultant IntegerVector and nil if successful, or an error if any of the checks fail.
func AddFloat64Vectors(vectors []Float64Vec, opts ...*concurrency.ConcurrencyOptions) (*Float64Vec, error) {
	vectors_len := len(vectors)
	if vectors_len == 0 {
		return nil, ErrEmptyVectorLength
	}

	// Get the size of the first vector
	size := vectors[0].Size()

	// Check whether all the vectors are of the same size
	for i := 1; i < vectors_len; i++ {
		if vectors[i].Size() != size {
			return nil, ErrInvalidVectorLength
		}
	}

	// Create a new vector with the same size and default value of 0
	resultant_vector := NewVector(Float64VecOptions{Size: size})

	

	// Add all vectors element-wise
	for _, vec := range vectors {
		for index := 0; index < size; index++ {
			resultant_vector.Data[index] += vec.Data[index]
		}
	}

	return resultant_vector, nil
}
