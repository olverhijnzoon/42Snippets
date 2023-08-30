CXX = g++
CXXFLAGS = -Wall -std=c++11
TARGET = electroweak_simulationX

all: $(TARGET)

$(TARGET): electroweak_simulation.cpp
	$(CXX) $(CXXFLAGS) -o $(TARGET) electroweak_simulation.cpp

run: $(TARGET)
	./$(TARGET)

clean:
	rm -f $(TARGET)
