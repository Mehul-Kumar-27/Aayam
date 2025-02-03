package matrix

import (
	assert "github.com/Mehul-Kumar-27/Aayam/utils"
)

type Float64Mat struct {
	Data [][]float64
}

// NewMatrix creates a new matrix based on the provided options
func NewMatrix(opts Float64MatOptions) *Float64Mat {
	if opts.Elements != nil {
		// If elements are provided, use them to create the matrix
		return &Float64Mat{Data: opts.Elements}
	}

	// If rows and cols are provided, create a matrix of that size
	if opts.Rows > 0 && opts.Cols > 0 {
		defaultValue := 0.0
		if opts.DefaultVal != nil {
			defaultValue = *opts.DefaultVal
		}

		// Initialize a 2D slice with rows and columns
		data := make([][]float64, opts.Rows)
		for i := range data {
			data[i] = make([]float64, opts.Cols)
			// Fill with the default value
			for j := range data[i] {
				data[i][j] = defaultValue
			}
		}
		return &Float64Mat{Data: data}
	}

	// If no options are provided, return an empty matrix
	return &Float64Mat{Data: [][]float64{}}
}

// Size returns the dimensions of the matrix (rows, columns)
func (mat *Float64Mat) Size() (int, int) {
	return len(mat.Data), len(mat.Data[0])
}

func (mat *Float64Mat) Rows() int {
	return len(mat.Data)
}

func (mat *Float64Mat) Columns() int {
	return len(mat.Data[0])
}

// ScalarMultiplication performs scalar multiplication on the matrix
func (mat *Float64Mat) ScalarMultiplication(scalar float64) {
	assert.AssertNotEqual(len(mat.Data), 0)
	for i := 0; i < len(mat.Data); i++ {
		for j := 0; j < len(mat.Data[i]); j++ {
			mat.Data[i][j] *= scalar
		}
	}
}

// GetVal returns the value at the specified row and column
func (mat *Float64Mat) GetVal(row, col int) float64 {
	assert.AssertRange[int](row, 0, len(mat.Data)-1)
	assert.AssertRange[int](col, 0, len(mat.Data[0])-1)
	return mat.Data[row][col]
}

// SetVal sets the value at the specified row and column
func (mat *Float64Mat) SetVal(row, col int, val float64) {
	assert.AssertRange[int](row, 0, len(mat.Data)-1)
	assert.AssertRange[int](col, 0, len(mat.Data[0])-1)
	mat.Data[row][col] = val
}

// DotProduct calculates the dot product of two matrices
func (mat *Float64Mat) DotProduct(another *Float64Mat) float64 {
	rows, cols := mat.Size()
	anotherRows, anotherCols := another.Size()
	assert.AssertEqual(rows, anotherRows)
	assert.AssertEqual(cols, anotherCols)

	var dotProduct float64
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			dotProduct += mat.GetVal(i, j) * another.GetVal(i, j)
		}
	}
	return dotProduct
}

// PushBack adds a row to the matrix
func (mat *Float64Mat) PushBack(row []float64) {
	mat.Data = append(mat.Data, row)
}

// PushFront adds a row to the front of the matrix
func (mat *Float64Mat) PushFront(row []float64) {
	mat.Data = append([][]float64{row}, mat.Data...)
}

func (mat *Float64Mat) GetColumn(col int) []float64 {
	assert.AssertRange[int](col, 0, len(mat.Data[0])-1)
	columnData := make([]float64, mat.Rows())

	for i := range mat.Data {
		columnData[i] = mat.Data[i][col]
	}
	return columnData
}

func (mat *Float64Mat) GetRow(row int) []float64 {
	assert.AssertRange[int](row, 0, mat.Rows())
	return mat.Data[row]
}
