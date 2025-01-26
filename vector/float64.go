package vector

import (
	assert "github.com/Mehul-Kumar-27/Aayam/utils"
)

type Float64Vec struct {
	Data []float64
}

func NewVector(opts Float64VecOptions) *Float64Vec {
	if opts.Elements != nil {
		// If elements are provided, use them to create the vector
		return &Float64Vec{Data: opts.Elements}
	}

	// If size is provided, create a vector of that size
	if opts.Size > 0 {
		defaultValue := 0.0
		if opts.DefaultVal != nil {
			defaultValue = *opts.DefaultVal
		}
		data := make([]float64, opts.Size)
		for i := range data {
			data[i] = defaultValue
		}
		return &Float64Vec{Data: data}
	}

	// If no options are provided, return an empty vector
	return &Float64Vec{Data: []float64{}}
}

func (vec *Float64Vec) Size() int {
	return len(vec.Data)
}

func (vec *Float64Vec) ScalarMultiplication(scalar float64) {
	var size int = vec.Size()
	assert.AssertNotEqual(size, 0)
	for i := 0; i < size; i++ {
		vec.Data[i] *= scalar
	}
}

func (vec *Float64Vec) GetVal(index int) float64 {
	assert.AssertRange[int](index, 0, vec.Size()-1)
	return vec.Data[index]
}

func (vec *Float64Vec) SetVal(index int, val float64) {
	assert.AssertRange[int](index, 0, vec.Size()-1)
	vec.Data[index] = val
}

func (vec *Float64Vec) DotProduct(another Float64Vec) float64 {
	assert.AssertEqual(vec.Size(), another.Size())
	var dot_product float64 = 0

	for i := 0; i < vec.Size(); i++ {
		dot_product += (vec.GetVal(i) * another.GetVal(i))
	}
	return dot_product
}

func (vec *Float64Vec) PushBack(element float64) {
	vec.Data = append(vec.Data, element)
}

func (vec *Float64Vec) PushFront(element float64) {
	vec.Data = append([]float64{element}, vec.Data...)
}
