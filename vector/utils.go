package vector

type BySize []Float64Vec

func (a BySize) Len() int           { return len(a) }
func (a BySize) Less(i, j int) bool { return a[i].Size() < a[j].Size() }
func (a BySize) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func UnmarshalFloat64ToVec(data []float64) Float64Vec {
	return *NewVector(Float64VecOptions{
		Elements: data,
	})
}
