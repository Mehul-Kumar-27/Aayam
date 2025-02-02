import json
import time
import numpy as np
import pytest

SMALL_DATA_SET_PATH = "data/vector_addition/small_data.jsonl"
SMALL_DATA_SET_SIZE = 11

NORMAL_DATA_SET_PATH = "data/vector_addition/normal_data.jsonl"
SMALL_DATA_SET_SIZE = 101

LARGE_DATA_SET_PATH = "data/vector_addition/large_data.jsonl"
SMALL_DATA_SET_SIZE = 1001

def read_line_data(reader):
    line = reader.readline()
    if not line:
        return None
    return unmarshal_data_set(line)

def unmarshal_data_set(line):
    temp = json.loads(line)
    return {
        "num_vectors": temp["num_vectors"],
        "vectors": [np.array(vec) for vec in temp["vectors"]],
        "sum": np.array(temp["sum"]),
    }

def add_float64_vectors(vectors):
    if not vectors:
        raise ValueError("Empty vector list")
    return np.sum(vectors, axis=0)

@pytest.mark.benchmark()
def test_vector_addition_small_dataset(benchmark):
    with open(SMALL_DATA_SET_PATH, "r") as file:
        reader = file
        def run_benchmark():
            total_time = 0  # Track time for the actual computation
            for _ in range(SMALL_DATA_SET_SIZE):
                # Read data (not included in the benchmark)
                vector_data = read_line_data(reader)
                if vector_data is None:
                    break

                # Measure only the time for vector addition and assertion
                start_time = time.perf_counter()
                sum_result = add_float64_vectors(vector_data["vectors"])
                assert np.allclose(sum_result, vector_data["sum"]), "Summation mismatch"
                end_time = time.perf_counter()

                # Accumulate the time for the computation
                total_time += (end_time - start_time)

            # Return the total time for the benchmark
            return total_time

        # Run the benchmark and report the total time
        total_time = benchmark(run_benchmark)
        print(f"Total computation time: {total_time:.6f} seconds")

@pytest.mark.benchmark()
def test_vector_addition_normal_dataset(benchmark):
    with open(NORMAL_DATA_SET_PATH, "r") as file:
        reader = file
        def run_benchmark():
            total_time = 0  # Track time for the actual computation
            for _ in range(SMALL_DATA_SET_SIZE):
                # Read data (not included in the benchmark)
                vector_data = read_line_data(reader)
                if vector_data is None:
                    break

                # Measure only the time for vector addition and assertion
                start_time = time.perf_counter()
                sum_result = add_float64_vectors(vector_data["vectors"])
                assert np.allclose(sum_result, vector_data["sum"]), "Summation mismatch"
                end_time = time.perf_counter()

                # Accumulate the time for the computation
                total_time += (end_time - start_time)

            # Return the total time for the benchmark
            return total_time

        # Run the benchmark and report the total time
        total_time = benchmark(run_benchmark)
        print(f"Total computation time: {total_time:.6f} seconds")

@pytest.mark.benchmark()
def test_vector_addition_large_dataset(benchmark):
    with open(LARGE_DATA_SET_PATH, "r") as file:
        reader = file
        def run_benchmark():
            total_time = 0  # Track time for the actual computation
            for _ in range(SMALL_DATA_SET_SIZE):
                # Read data (not included in the benchmark)
                vector_data = read_line_data(reader)
                if vector_data is None:
                    break

                # Measure only the time for vector addition and assertion
                start_time = time.perf_counter()
                sum_result = add_float64_vectors(vector_data["vectors"])
                assert np.allclose(sum_result, vector_data["sum"]), "Summation mismatch"
                end_time = time.perf_counter()

                # Accumulate the time for the computation
                total_time += (end_time - start_time)

            # Return the total time for the benchmark
            return total_time

        # Run the benchmark and report the total time
        total_time = benchmark(run_benchmark)
        print(f"Total computation time: {total_time:.6f} seconds")
