# Compiler and flags
CXX = g++
CXXFLAGS = -Wall -std=c++20

# Source files and objects
SOURCES = electroweak_simulation.cpp particle.cpp boson.cpp fermion.cpp higgs_field.cpp
OBJECTS = $(SOURCES:.cpp=.o)
TARGET = electroweak_simulationX

# Default target
all: $(TARGET)

$(TARGET): $(OBJECTS)
	$(CXX) $(CXXFLAGS) -o $@ $^

# Generic rule for compiling source files
%.o: %.cpp
	$(CXX) $(CXXFLAGS) -c $< -o $@

# Run target
run: $(TARGET)
	./$(TARGET)

# Clean target
clean:
	rm -f $(OBJECTS) $(TARGET)

.PHONY: all run clean
