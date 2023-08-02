# List: Ordered, mutable collection
fruits = ['apple', 'banana', 'cherry']
print("List:", fruits)

# Tuple: Ordered, immutable collection
coordinates = (4, 5)
print("Tuple:", coordinates)

# Set: Unordered collection with no duplicate elements
unique_numbers = {1, 2, 3, 2, 1}
print("Set:", unique_numbers)

# Dictionary: Unordered collection of key-value pairs
student = {
    'name': 'John',
    'age': 20,
    'courses': ['math', 'physics']
}
print("Dictionary:", student)

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
print("Nested Data Structures:", nested_structure)

# List Comprehensions
squares = [x**2 for x in range(10)]
print("List Comprehensions:", squares)

# Deque
from collections import deque
d = deque([1, 2, 3])
print("Deque:", d)

# Counter
from collections import Counter
counts = Counter(['a', 'b', 'a', 'c', 'b', 'b'])
print("Counter:", counts)

# NamedTuple
from collections import namedtuple
Point = namedtuple('Point', ['x', 'y'])
p = Point(11, y=22)
print("NamedTuple:", p)

# DefaultDict
from collections import defaultdict
dd = defaultdict(int)
dd['key'] += 1
print("DefaultDict:", dd)

# Heapq
import heapq
heap = [3, 1, 4, 1, 5, 9, 2]
heapq.heapify(heap)
print("Heap:", heap)

# OrderedDict
from collections import OrderedDict
od = OrderedDict.fromkeys('abcde')
print("OrderedDict:", od)

# ChainMap
from collections import ChainMap
baseline = {'a': 1, 'b': 2}
adjustments = {'b': 3, 'c': 4}
combined = ChainMap(adjustments, baseline)
print("ChainMap:", combined)

# Frozenset
fs = frozenset([1, 2, 3, 2, 1])
print("Frozenset:", fs)

# Array
from array import array
arr = array('I', [1, 2, 3, 4])
print("Array:", arr)
