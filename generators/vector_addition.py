import json
import os
import numpy as np

def generate_vector_data(file_name, num_entries=1000, vector_size=5000, value_range=(-100, 100)):
    """
    Generates a JSON Lines file containing vector data with a random number of vectors and their sums using NumPy.

    :param file_name: Path to the output JSON Lines file
    :param num_entries: Number of entries in the JSON Lines file
    :param vector_size: Size of each vector
    :param value_range: Range of values for vector elements (min, max)
    """
    # Ensure the directory exists
    os.makedirs(os.path.dirname(file_name), exist_ok=True)

    # Get the absolute path
    absolute_path = os.path.abspath(file_name)

    # Open the file in write mode
    with open(file_name, "w") as f:
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
            # Create a dictionary for the current entry
            entry = {"num_vectors": num_vectors, "vectors": vectors, "sum": sum_vector}
            # Write the entry as a JSON line
            f.write(json.dumps(entry) + "\n")

    # Print the exact path where the data is dumped
    print(f"Data saved at: {absolute_path}")


# Generate different sizes of data
generate_vector_data("data/vector_addition/small_data.jsonl", num_entries=10, vector_size=50, value_range=(-1000, 1000))
generate_vector_data("data/vector_addition/normal_data.jsonl", num_entries=100, vector_size=500, value_range=(-1e6, 1e6))
generate_vector_data("data/vector_addition/large_data.jsonl", num_entries=1000, vector_size=5000, value_range=(-1e9, 1e9))
