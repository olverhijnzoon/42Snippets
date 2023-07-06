#!/bin/bash

# Run sar command
echo "Collecting system activity information"
sar 1 1
echo "--------------------------------------------"

# Print CPU usage
echo "Printing CPU usage"
sar -u 1 1
echo "--------------------------------------------"

# Print I/O statistics
echo "Printing I/O statistics"
sar -b 1 1
echo "--------------------------------------------"

# Print memory usage statistics
echo "Printing memory usage statistics"
sar -r 1 1
echo "--------------------------------------------"

# List running processes
echo "Listing running processes"
ps -eo pid,user,pcpu,pmem,vsz,rss,comm | head -10
echo "--------------------------------------------"