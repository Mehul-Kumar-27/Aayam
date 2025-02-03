package matrix

import (
	"reflect"
	"testing"

	"github.com/Mehul-Kumar-27/Aayam/concurrency"
)

func TestAddMatrix(t *testing.T) {
	tests := []struct {
		name           string
		matrices       []Float64Mat
		expectedResult *Float64Mat
		expectedError  error
		isConcurrent   bool
	}{
		{
			name: "Add two matrices",
			matrices: []Float64Mat{
				*NewMatrix(Float64MatOptions{Rows: 2, Cols: 2, Elements: [][]float64{{1, 2}, {3, 4}}}),
				*NewMatrix(Float64MatOptions{Rows: 2, Cols: 2, Elements: [][]float64{{5, 6}, {7, 8}}}),
			},
			expectedResult: NewMatrix(Float64MatOptions{Rows: 2, Cols: 2, Elements: [][]float64{{6, 8}, {10, 12}}}),
			expectedError:  nil,
			isConcurrent:   false,
		},
		{
			name: "Add three matrices",
			matrices: []Float64Mat{
				*NewMatrix(Float64MatOptions{Rows: 2, Cols: 2, Elements: [][]float64{{1, 1}, {1, 1}}}),
				*NewMatrix(Float64MatOptions{Rows: 2, Cols: 2, Elements: [][]float64{{2, 2}, {2, 2}}}),
				*NewMatrix(Float64MatOptions{Rows: 2, Cols: 2, Elements: [][]float64{{3, 3}, {3, 3}}}),
			},
			expectedResult: NewMatrix(Float64MatOptions{Rows: 2, Cols: 2, Elements: [][]float64{{6, 6}, {6, 6}}}),
			expectedError:  nil,
			isConcurrent:   false,
		},
		{
			name: "Add matrices with different dimensions",
			matrices: []Float64Mat{
				*NewMatrix(Float64MatOptions{Rows: 2, Cols: 2, Elements: [][]float64{{1, 2}, {3, 4}}}),
				*NewMatrix(Float64MatOptions{Rows: 3, Cols: 2, Elements: [][]float64{{5, 6}, {7, 8}, {9, 10}}}),
			},
			expectedResult: nil,
			expectedError:  ErrDimensionMismatch,
			isConcurrent:   false,
		},
		{
			name:           "Add empty matrices slice",
			matrices:       []Float64Mat{},
			expectedResult: nil,
			expectedError:  ErrEmptySliceForMatrixAddition,
			isConcurrent:   false,
		},
		{
			name: "Add zero matrices",
			matrices: []Float64Mat{
				*NewMatrix(Float64MatOptions{Rows: 2, Cols: 2, Elements: [][]float64{{0, 0}, {0, 0}}}),
				*NewMatrix(Float64MatOptions{Rows: 2, Cols: 2, Elements: [][]float64{{0, 0}, {0, 0}}}),
			},
			expectedResult: NewMatrix(Float64MatOptions{Rows: 2, Cols: 2, Elements: [][]float64{{0, 0}, {0, 0}}}),
			expectedError:  nil,
			isConcurrent:   false,
		},
		{
			name: "Add two matrices concurrently",
			matrices: []Float64Mat{
				*NewMatrix(Float64MatOptions{Rows: 3, Cols: 3, Elements: [][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}),
				*NewMatrix(Float64MatOptions{Rows: 3, Cols: 3, Elements: [][]float64{{9, 8, 7}, {6, 5, 4}, {3, 2, 1}}}),
			},
			expectedResult: NewMatrix(Float64MatOptions{Rows: 3, Cols: 3, Elements: [][]float64{{10, 10, 10}, {10, 10, 10}, {10, 10, 10}}}),
			expectedError:  nil,
			isConcurrent:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := AddMatrix(tt.matrices, &concurrency.ConcurrencyOptions{Enabled: tt.isConcurrent, Batch_Size: 10})
			if (err != nil && tt.expectedError == nil) || (err == nil && tt.expectedError != nil) || (err != nil && tt.expectedError != nil && err.Error() != tt.expectedError.Error()) {
				t.Errorf("AddMatrix() error = %v, wantErr %v", err, tt.expectedError)
				return
			}
			if !reflect.DeepEqual(result, tt.expectedResult) {
				t.Errorf("AddMatrix() = %v, want %v", result, tt.expectedResult)
			}
		})
	}
}
