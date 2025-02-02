package vectors

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"

	vec "github.com/Mehul-Kumar-27/Aayam/vector"
)

// Represent the Testing data for vector addition
type VectorData struct {
	// Number of vectors in the dataset
	NumVectors int `json:"num_vectors"`
	// List of vectors
	Vectors []vec.Float64Vec `json:"vectors"`
	// Sum of these vectors
	Sum vec.Float64Vec `json:"sum"`
}

func UnmarshalDataSet(line []byte) (*VectorData, error) {
	type Alias struct {
		NumVectors int         `json:"num_vectors"`
		Vectors    [][]float64 `json:"vectors"`
		Sum        []float64   `json:"sum"`
	}
	var temp Alias
	if err := json.Unmarshal(line, &temp); err != nil {
		return nil, err
	}
	vd := &VectorData{
		NumVectors: temp.NumVectors,
		Sum:        vec.UnmarshalFloat64ToVec(temp.Sum),
		Vectors:    make([]vec.Float64Vec, len(temp.Vectors)),
	}

	for i, vector := range temp.Vectors {
		vd.Vectors[i] = vec.UnmarshalFloat64ToVec(vector)
	}
	return vd, nil
}

func ReadVectorData(filename string) ([]VectorData, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 0

	var vectorData []VectorData

	for scanner.Scan() {
		lineNumber++
		line_byte := scanner.Bytes()

		var entry VectorData
		if err := json.Unmarshal(line_byte, &entry); err != nil {
			return nil, fmt.Errorf("error parsing the jsonl data %d: %v", lineNumber, err)
		}
		vectorData = append(vectorData, entry)
	}
	return vectorData, nil
}

func getLineData(scanner *bufio.Scanner, line int) (*VectorData, error) {
	var lineNumber int = 1
	for scanner.Scan() {
		if lineNumber == line {
			lineByte := scanner.Bytes()

			return UnmarshalDataSet(lineByte)
		}
		lineNumber++
	}
	return nil, fmt.Errorf("no data to get")

}

func getLineDataUsingReader(reader *bufio.Reader) (*VectorData, error) {

	lineByte, err := reader.ReadBytes('\n')
	if err != nil {
		if err == io.EOF {
			return nil, err
		}
		return nil, fmt.Errorf("error reading the line :%v", err)
	}
	return UnmarshalDataSet(lineByte)

}
