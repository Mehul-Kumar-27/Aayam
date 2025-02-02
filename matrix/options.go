package matrix

type Float64MatOptions struct {
	Rows       int         // Number of rows
	Cols       int         // Number of columns
	DefaultVal *float64    // Default value for each element
	Elements   [][]float64 // Predefined elements
}
