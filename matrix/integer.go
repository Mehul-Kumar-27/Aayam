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

func (mat *IntegerMatrix) GetVal(row, column int) int {
	assert.AssertRange[int](row, 0, mat.Rows-1)
	assert.AssertRange[int](column, 0, mat.Columns-1)

	return mat.Data[row].GetVal(column)
}

func (mat *IntegerMatrix) GetRow(row int) vec.IntegerVector {
	assert.AssertRange[int](row, 0, mat.Rows-1)
	return mat.Data[row]
}

func (mat *IntegerMatrix) ScalarMultiplication(scalar int) {
	assert.AssertNotEqual(mat.Rows, 0)
	assert.AssertNotEqual(mat.Columns, 0)

	for row := 0; row < mat.Rows; row++ {
		mat.Data[row].ScalarMultiplication(scalar)
	}
}
