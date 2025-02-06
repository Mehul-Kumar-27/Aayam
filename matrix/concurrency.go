package matrix

import "sync"

func addMatrixRowParallel(mats []Float64Mat, resultantMat *Float64Mat, numWorkers int) error {
	rows := resultantMat.Rows()
	cols := resultantMat.Columns()
	var wg sync.WaitGroup

	// Create a channel to distribute row indices to workers.
	rowChan := make(chan int, rows)

	// Start worker goroutines.
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Each worker processes rows sent through the channel.
			for r := range rowChan {
				// For each column in the row, compute the sum over all matrices.
				for c := 0; c < cols; c++ {
					var sum float64
					for _, mat := range mats {
						sum += mat.Data[r][c]
					}
					resultantMat.Data[r][c] = sum
				}
			}
		}()
	}

	for r := 0; r < rows; r++ {
		rowChan <- r
	}
	close(rowChan)

	wg.Wait()
	return nil
}
