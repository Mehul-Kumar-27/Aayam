package vectors

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"testing"

	vec "github.com/Mehul-Kumar-27/Aayam/vector"
)

const smallDataSetSize int = 11
const normalDataSetSize int = 101
const largeDataSetSize int = 1001

func BenchmarkVectorAdditionSmallDataSet(b *testing.B) {
	dataSetPath := "../../../data/vector_addition/small_data.jsonl"
	cwd, err := os.Getwd()
	if err != nil {
		b.Fatalf("error getting current working directory: %v", err)
	}
	dataSetPath = filepath.Join(cwd, dataSetPath)
	file, err := os.Open(dataSetPath)
	if err != nil {
		b.Fatalf("error opening the dataset path %v", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	for line := 1; line <= smallDataSetSize; line++ {
		b.StopTimer()
		vectorData, err := getLineDataUsingReader(reader)
		if err == io.EOF {
			break
		}
		b.StartTimer()
		if err != nil {
			b.Fatalf("%v", err)
		}
		_, err = vec.AddFloat64Vectors(vectorData.Vectors)
		if err != nil {
			b.Fatalf("error occurred while adding vectors :%v", err)
		}
	}

}

func BenchmarkVectorAdditionNormalDataSet(b *testing.B) {
	dataSetPath := "../../../data/vector_addition/normal_data.jsonl"
	cwd, err := os.Getwd()
	if err != nil {
		b.Fatalf("error getting current working directory: %v", err)
	}
	dataSetPath = filepath.Join(cwd, dataSetPath)
	file, err := os.Open(dataSetPath)
	if err != nil {
		b.Fatalf("error opening the dataset path %v", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	for line := 1; line <= normalDataSetSize; line++ {
		b.StopTimer()
		vectorData, err := getLineDataUsingReader(reader)
		if err == io.EOF {
			break
		}
		b.StartTimer()
		if err != nil {
			b.Fatalf("%v", err)
		}
		_, err = vec.AddFloat64Vectors(vectorData.Vectors)
		if err != nil {
			b.Fatalf("error occurred while adding vectors :%v", err)
		}
	}

}

func BenchmarkVectorAdditionLargeDataSet(b *testing.B) {
	dataSetPath := "../../../data/vector_addition/large_data.jsonl"
	cwd, err := os.Getwd()
	if err != nil {
		b.Fatalf("error getting current working directory: %v", err)
	}
	dataSetPath = filepath.Join(cwd, dataSetPath)
	file, err := os.Open(dataSetPath)
	if err != nil {
		b.Fatalf("error opening the dataset path %v", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	for line := 1; line <= largeDataSetSize; line++ {
		b.StopTimer()
		vectorData, err := getLineDataUsingReader(reader)
		if err == io.EOF {
			break
		}
		b.StartTimer()
		if err != nil {
			b.Fatalf("%v", err)
		}
		_, err = vec.AddFloat64Vectors(vectorData.Vectors)
		if err != nil {
			b.Fatalf("error occurred while adding vectors :%v", err)
		}
	}

}
