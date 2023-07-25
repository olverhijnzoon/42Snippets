# Makefile
CC = gcc
CFLAGS = -Wall
TARGET = network_namespace

all: $(TARGET)

$(TARGET): $(TARGET).c
	$(CC) $(CFLAGS) -o $(TARGET) $(TARGET).c

clean:
	rm -f $(TARGET)

# The --privileged flag is required to allow the Docker container to create new namespaces.
it_to_docker:
	docker build -t $(TARGET) .
	docker run -it --rm --privileged $(TARGET)
	# docker run -it --rm --privileged $(TARGET) bash

