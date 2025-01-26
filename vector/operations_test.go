package vector

import (
	"reflect"
	"testing"

	"github.com/Mehul-Kumar-27/Aayam/concurrency"
)

func TestAddFloat64Vectors(t *testing.T) {
	tests := []struct {
		name           string
		vectors        []Float64Vec
		expectedResult *Float64Vec
		expectedError  error
		isConcurrent   bool
	}{
		{
			name: "Add two vectors",
			vectors: []Float64Vec{
				*NewVector(Float64VecOptions{Size: 3, Elements: []float64{1, 2, 3}}),
				*NewVector(Float64VecOptions{Size: 3, Elements: []float64{4, 5, 6}}),
			},
			expectedResult: NewVector(Float64VecOptions{Size: 3, Elements: []float64{5, 7, 9}}),
			expectedError:  nil,
			isConcurrent:   false,
		},
		{
			name: "Add three vectors",
			vectors: []Float64Vec{
				*NewVector(Float64VecOptions{Size: 3, Elements: []float64{1, 2, 3}}),
				*NewVector(Float64VecOptions{Size: 3, Elements: []float64{4, 5, 6}}),
				*NewVector(Float64VecOptions{Size: 3, Elements: []float64{7, 8, 9}}),
			},
			expectedResult: NewVector(Float64VecOptions{Size: 3, Elements: []float64{12, 15, 18}}),
			expectedError:  nil,
			isConcurrent:   false,
		},
		{
			name: "Add vectors with different sizes",
			vectors: []Float64Vec{
				*NewVector(Float64VecOptions{Size: 3, Elements: []float64{1, 2, 3}}),
				*NewVector(Float64VecOptions{Size: 4, Elements: []float64{4, 5, 6, 7}}),
			},
			expectedResult: NewVector(Float64VecOptions{Size: 3, Elements: []float64{5, 7, 9, 7}}),
			expectedError:  nil,
		},
		{
			name: "Add empty vectors",
			vectors: []Float64Vec{
				*NewVector(Float64VecOptions{}),
				*NewVector(Float64VecOptions{}),
			},
			expectedResult: nil,
			expectedError:  ErrEmptyVectorLength,
			isConcurrent:   false,
		},
		{
			name: "Add Zero Vectors",
			vectors: []Float64Vec{
				*NewVector(Float64VecOptions{Elements: []float64{0}}),
				*NewVector(Float64VecOptions{Elements: []float64{0}}),
			},
			expectedResult: NewVector(Float64VecOptions{Size: 1, Elements: []float64{0}}),
			expectedError:  nil,
			isConcurrent:   false,
		},
		{
			name: "Add two vectors Concurrently",
			vectors: []Float64Vec{
				*NewVector(Float64VecOptions{
					Size: 20,
					Elements: []float64{
						1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
						11, 12, 13, 14, 15, 16, 17, 18, 19,
					},
				}),
				*NewVector(Float64VecOptions{
					Size: 20,
					Elements: []float64{
						20, 19, 18, 17, 16, 15, 14, 13, 12, 11,
						10, 9, 8, 7, 6, 5, 4, 3, 2, 1,
					},
				}),
			},
			expectedResult: NewVector(Float64VecOptions{
				Size: 20,
				Elements: []float64{
					21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
					21, 21, 21, 21, 21, 21, 21, 21, 21, 1,
				},
			}),
			expectedError: nil,
			isConcurrent:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := AddFloat64Vectors(tt.vectors, &concurrency.ConcurrencyOptions{Enabled: tt.isConcurrent, Batch_Size: 10})
			if (err != nil && tt.expectedError == nil) || (err == nil && tt.expectedError != nil) || (err != nil && tt.expectedError != nil && err.Error() != tt.expectedError.Error()) {
				t.Errorf("AddFloat64Vectors() error = %v, wantErr %v", err, tt.expectedError)
				return
			}
			if !reflect.DeepEqual(result, tt.expectedResult) {
				t.Errorf("AddFloat64Vectors() = %v, want %v", result, tt.expectedResult)
			}
		})
	}
}
