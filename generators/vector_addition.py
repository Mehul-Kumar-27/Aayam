import json
import os
import numpy as np

def generate_vector_data(file_name, num_entries=1000, vector_size=5000, value_range=(-100, 100)):
    """
    Generates a JSON file containing vector data with a random number of vectors and their sums using NumPy.

    :param file_name: Path to the output JSON file
    :param num_entries: Number of entries in the JSON file
    :param vector_size: Size of each vector
    :param value_range: Range of values for vector elements (min, max)
    """
    # Ensure the directory exists
    os.makedirs(os.path.dirname(file_name), exist_ok=True)

    # Generate the data
    data = []
    for _ in range(num_entries):
        # Randomly determine the number of vectors (between 1 and 100)
        num_vectors = np.random.randint(1, 101)
        # Generate the vectors
        vectors = [
            np.random.uniform(value_range[0], value_range[1], size=vector_size).tolist()
            for _ in range(num_vectors)
        ]
        # Calculate the sum of all vectors
        sum_vector = np.sum(vectors, axis=0).tolist()
        data.append({"num_vectors": num_vectors, "vectors": vectors, "sum": sum_vector})

    # Get the absolute path
    absolute_path = os.path.abspath(file_name)

    # Write data to the file
    with open(file_name, "w") as f:
        json.dump(data, f, indent=4)

    # Print the exact path where the data is dumped
    print(f"Data saved at: {absolute_path}")


# Generate different sizes of data
generate_vector_data("data/vector_addition/small_data.json", num_entries=10, vector_size=50, value_range=(-1000, 1000))
generate_vector_data("data/vector_addition/normal_data.json", num_entries=100, vector_size=500, value_range=(-1e6, 1e6))
generate_vector_data("data/vector_addition/large_data.json", num_entries=1000, vector_size=5000, value_range=(-1e9, 1e9))
