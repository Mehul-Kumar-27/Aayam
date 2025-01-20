package vector

import (
	assert "github.com/Mehul-Kumar-27/Aayam/utils"
)

type IntegerVector struct {
	Data []int
}

func NewIntegerVector(elements ...int) Vector {
	return &IntegerVector{
		Data: elements,
	}
}

func (vec *IntegerVector) Size() int {
	return len(vec.Data)
}

func (vec *IntegerVector) ScalarMultiplication(scalar int) {
	var size int = vec.Size()
	assert.AssertNotEqual(size, 0)
	for i := 0; i < size; i++ {
		vec.Data[i] *= scalar
	}
}

func (vec *IntegerVector) GetVal(index int) int {
	assert.AssertRange[int](index, 0, vec.Size()-1)
	return vec.Data[index]
}

func (vec *IntegerVector) DotProduct(another IntegerVector) int {
	assert.AssertEqual(vec.Size(), another.Size())
	var dot_product int = 0

	for i := 0; i < vec.Size(); i++ {
		dot_product += (vec.GetVal(i) * another.GetVal(i))
	}
	return dot_product
}

func (vec *IntegerVector) PushBack(element int) {
	vec.Data = append(vec.Data, element)
}

func (vec *IntegerVector) PushFront(element int) {
	vec.Data = append([]int{element}, vec.Data...)
}
