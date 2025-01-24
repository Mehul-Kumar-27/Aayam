package matrix

import (
	assert "github.com/Mehul-Kumar-27/Aayam/utils"
	vec "github.com/Mehul-Kumar-27/Aayam/vector"
)

type IntegerMatrix struct {
	Data    []vec.IntegerVector
	Rows    int
	Columns int
}

func NewIntegerMatrix(elements ...vec.IntegerVector) *IntegerMatrix {
	rows := len(elements)
	columns := 0
	if rows > 0 {
		columns = elements[0].Size()
	}
	integer_matrix := &IntegerMatrix{
		Data:    elements,
		Rows:    rows,
		Columns: columns,
	}
	return integer_matrix
}

func NewIntegerMatrixWithDimensions(row, column int, default_value *float64) *IntegerMatrix {
	var default_val float64 = 0
	if default_value != nil {
		default_val = *default_value
	}
	var rows = make([]vec.IntegerVector, row)
	for i := 0; i < row; i++ {
		rows = append(rows, *vec.NewIntegerVectorWithSize(column, &default_val))
	}
	return NewIntegerMatrix(rows...)
}

func (mat *IntegerMatrix) GetVal(row, column int) float64 {
	assert.AssertRange[int](row, 0, mat.Rows-1)
	assert.AssertRange[int](column, 0, mat.Columns-1)

	return mat.Data[row].GetVal(column)
}

func (mat *IntegerMatrix) GetRow(row int) vec.IntegerVector {
	assert.AssertRange[int](row, 0, mat.Rows-1)
	return mat.Data[row]
}

func (mat *IntegerMatrix) GetColumn(col int) vec.IntegerVector {
	assert.AssertRange[int](col, 0, mat.Columns-1)
	columnVector := vec.NewIntegerVector()
	for row := 0; row < mat.Rows; row++ {
		columnVector.PushBack(mat.GetVal(row, col))
	}
	return *columnVector
}

func (mat *IntegerMatrix) ScalarMultiplication(scalar float64) {
	assert.AssertNotEqual(mat.Rows, 0)
	assert.AssertNotEqual(mat.Columns, 0)

	for row := 0; row < mat.Rows; row++ {
		mat.Data[row].ScalarMultiplication(scalar)
	}
}

func (mat *IntegerMatrix) SetVal(row, column int, val float64) {
	assert.AssertRange[int](row, 0, mat.Rows-1)
	assert.AssertRange[int](column, 0, mat.Columns-1)
	mat.Data[row].SetVal(column, val)
}

// Multiplies two matrices the argument matrix is present on the right of the matrix
// It does the multiplication by representing the multiplication as the linear combination of columns of the two matrices
func (mat *IntegerMatrix) SimpleMultiplication(right IntegerMatrix) (*IntegerMatrix, error) {
	if !assert.AssertEqual(mat.Columns, right.Rows) {
		return nil, ErrDimensionMismatch
	}
	result := NewIntegerMatrixWithDimensions(mat.Rows, right.Columns, nil)
	var result_col int = result.Columns

	for result_column := 0 ; result_column < result_col; result_column++{
		var columns_to_add []vec.IntegerVector = make([]vec.IntegerVector, mat.Columns)
		for right_col := 0 ; right_col < mat.Columns; right_col++{
			col := mat.GetColumn(right_col)
			col.ScalarMultiplication(right.GetVal(right_col, result_column))
			columns_to_add = append(columns_to_add, col)
		}
		
	}
	return result, nil
}
