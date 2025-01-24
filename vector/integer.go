package vector

import (
	assert "github.com/Mehul-Kumar-27/Aayam/utils"
)

type IntegerVector struct {
	Data []float64
}

func NewIntegerVector(elements ...float64) *IntegerVector {
	return &IntegerVector{
		Data: elements,
	}
}

func NewIntegerVectorWithSize(size int, default_val *float64) *IntegerVector {
	var default_value float64 = 0
	if default_val != nil {
		default_value = *default_val
	}
	var data []float64 = make([]float64, size)
	for i := 0; i < size; i++ {
		data[i] = default_value
	}
	return NewIntegerVector(data...)
}

func (vec *IntegerVector) Size() int {
	return len(vec.Data)
}

func (vec *IntegerVector) ScalarMultiplication(scalar float64) {
	var size int = vec.Size()
	assert.AssertNotEqual(size, 0)
	for i := 0; i < size; i++ {
		vec.Data[i] *= scalar
	}
}

func (vec *IntegerVector) GetVal(index int) float64 {
	assert.AssertRange[int](index, 0, vec.Size()-1)
	return vec.Data[index]
}

func (vec *IntegerVector) SetVal(index int, val float64) {
	assert.AssertRange[int](index, 0, vec.Size()-1)
	vec.Data[index] = val
}

func (vec *IntegerVector) DotProduct(another IntegerVector) float64 {
	assert.AssertEqual(vec.Size(), another.Size())
	var dot_product float64 = 0

	for i := 0; i < vec.Size(); i++ {
		dot_product += (vec.GetVal(i) * another.GetVal(i))
	}
	return dot_product
}

func (vec *IntegerVector) PushBack(element float64) {
	vec.Data = append(vec.Data, element)
}

func (vec *IntegerVector) PushFront(element float64) {
	vec.Data = append([]float64{element}, vec.Data...)
}
