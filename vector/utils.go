package vector

import (
	assert "github.com/Mehul-Kumar-27/Aayam/utils"
)

func AddVectors(vectors ...IntegerVector) (*IntegerVector, error) {
	if len(vectors) == 0 {
		return nil,  ErrEmptyVectorLength
	}
	initial_vector := vectors[0]
	var size int = initial_vector.Size()

	// Check whether all the vectors are of the same size
	for i := 0 ; i < len(vectors); i++{
		if !assert.AssertEqual(vectors[i].Size(), size){
			return nil, ErrInvalidVectorLength
		}
	}

	// Iterate over the vectors present
	for i := 1; i < len(vectors); i++{
		// Iterate over each index of the vector and add it to the initial vector
		for j := 0; j < size; j++ {
			initial_vector[j] += vectors[i].GetVal(j)
		}
	}
}
