#!/bin/bash

interval=1
count=1

# Run sar command
echo "Collecting system activity information"
sar $interval $count
echo "--------------------------------------------"

# Print CPU usage
echo "Printing CPU usage"
sar -u $interval $count
echo "--------------------------------------------"

# Print I/O statistics
echo "Printing I/O statistics"
sar -b $interval $count
echo "--------------------------------------------"

# Print memory usage statistics
echo "Printing memory usage statistics"
sar -r $interval $count
echo "--------------------------------------------"