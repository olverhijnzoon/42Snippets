# List: Ordered, mutable collection
fruits = ['apple', 'banana', 'cherry']

# Tuple: Ordered, immutable collection
coordinates = (4, 5)

# Set: Unordered collection with no duplicate elements
unique_numbers = {1, 2, 3, 2, 1}

# Dictionary: Unordered collection of key-value pairs
student = {
    'name': 'John',
    'age': 20,
    'courses': ['math', 'physics']
}

# Custom Class
class Car:
    def __init__(self, make, model):
        self.make = make
        self.model = model

    def display_info(self):
        print(f"This car is a {self.make} {self.model}")

# Using the custom class
my_car = Car('Toyota', 'Camry')
my_car.display_info()

# Nested Data Structures
nested_structure = {
    'list_inside_dict': [1, 2, 3],
    'tuple_inside_dict': (4, 5),
    'dict_inside_dict': {'key': 'value'},
    'set_inside_dict': {6, 7, 8}
}

# Accessing elements
print(fruits[0])  # Output: apple
print(coordinates[1])  # Output: 5
print(unique_numbers)  # Output: {1, 2, 3}
print(student['name'])  # Output: John
print(nested_structure['list_inside_dict'][2])  # Output: 3
