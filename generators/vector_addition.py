import json
import os
import numpy as np

def generate_vector_data(file_name, num_vectors=1000, vector_size=5000, value_range=(-100, 100)):
    """
    Generates a JSON file containing vector data with their sums using NumPy.

    :param file_name: Path to the output JSON file
    :param num_vectors: Number of vector pairs to generate
    :param vector_size: Size of each vector
    :param value_range: Range of values for vector elements (min, max)
    """
    # Ensure the directory exists
    os.makedirs(os.path.dirname(file_name), exist_ok=True)

    # Generate the data
    data = []
    for _ in range(num_vectors):
        vec1 = np.random.randint(value_range[0], value_range[1], size=vector_size).tolist()
        vec2 = np.random.randint(value_range[0], value_range[1], size=vector_size).tolist()
        result = (np.array(vec1) + np.array(vec2)).tolist()
        data.append({"vector1": vec1, "vector2": vec2, "sum": result})

    # Get the absolute path
    absolute_path = os.path.abspath(file_name)

    # Write data to the file
    with open(file_name, "w") as f:
        json.dump(data, f, indent=4)

    # Print the exact path where the data is dumped
    print(f"Data saved at: {absolute_path}")




generate_vector_data("data/vector_addition/small_data.json", num_vectors=10, vector_size=50)
generate_vector_data("data/vector_addition/normal_data.json", num_vectors=100, vector_size=500)
generate_vector_data("data/vector_addition/large_data.json", num_vectors=1000, vector_size=5000)
