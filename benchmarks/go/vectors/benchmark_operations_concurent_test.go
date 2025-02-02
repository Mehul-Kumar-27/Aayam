package vectors

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/Mehul-Kumar-27/Aayam/concurrency"
	vec "github.com/Mehul-Kumar-27/Aayam/vector"
)

func BenchmarkVectorAdditionSmallDataSetConcurrently(b *testing.B) {
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
		concurrencyOpts := concurrency.ConcurrencyOptions{
			Enabled: true,
			Batch_Size: len(vectorData.Vectors) / 20,
		}
		if err != nil {
			b.Fatalf("%v", err)
		}
		_, err = vec.AddFloat64Vectors(vectorData.Vectors, &concurrencyOpts)
		if err != nil {
			b.Fatalf("error occurred while adding vectors :%v", err)
		}
	}

}

func BenchmarkVectorAdditionNormalDataSetConcurrently(b *testing.B) {
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
		concurrencyOpts := concurrency.ConcurrencyOptions{
			Enabled: true,
			Batch_Size: len(vectorData.Vectors) / 20,
		}
		_, err = vec.AddFloat64Vectors(vectorData.Vectors, &concurrencyOpts)
		if err != nil {
			b.Fatalf("error occurred while adding vectors :%v", err)
		}
	}

}

func BenchmarkVectorAdditionLargeDataSetConcurrently(b *testing.B) {
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
		concurrencyOpts := concurrency.ConcurrencyOptions{
			Enabled: true,
			Batch_Size: len(vectorData.Vectors) / 20,
		}
		if err != nil {
			b.Fatalf("%v", err)
		}
		_, err = vec.AddFloat64Vectors(vectorData.Vectors, &concurrencyOpts)
		if err != nil {
			b.Fatalf("error occurred while adding vectors :%v", err)
		}
	}

}
