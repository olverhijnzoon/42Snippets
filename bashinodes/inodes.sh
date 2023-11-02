#!/bin/bash

# Function to check if command exists
function command_exists() {
  type "$1" &> /dev/null ;
}

# Check if ls and stat commands exist
if ! command_exists ls || ! command_exists stat; then
  echo "ls and/or stat not found. Please ensure they are installed and try again."
  exit 1
fi

# Check if a directory or file to check was supplied
if [ -z "$1" ]; then
  echo "Usage: $0 <directory_or_file>"
  exit 1
fi

directory_or_file="$1"

# Check if the provided argument is a directory or a file
if [ -d "$directory_or_file" ]; then
  mode="directory"
elif [ -f "$directory_or_file" ]; then
  mode="file"
else
  echo "The provided argument is neither a file nor a directory."
  exit 1
fi

# Print debug information
echo "Debug: inode details for $mode '$directory_or_file'"

# Define a function to print information in a formatted way
function print_inode_details() {
  local item=$1
  # Use the stat command to retrieve inode information
  local inode_info=$(stat -c "Inode: %i  Size: %s bytes  Permissions: %A  Owner: %U  File Type: %F" "$item")
  echo -e "\e[34m$inode_info\e[0m"
}

# "\033[H\033[2J" - Clears the terminal screen.
echo -e "\033[H\033[2J\033[34m===== Inode Details =====\e[0m"

# If it's a directory, list files and show inode details; if it's a file, just show the inode details
if [ "$mode" == "directory" ]; then
  for file in "$directory_or_file"/*; do
    [ -e "$file" ] || continue # Ensuring the file exists before getting its details
    print_inode_details "$file"
  done
else
  print_inode_details "$directory_or_file"
fi
