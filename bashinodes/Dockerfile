# Start with the Ubuntu base image
FROM ubuntu:latest

# Update the package list and install any required packages if necessary
# (In this case, coreutils is usually pre-installed in ubuntu, which includes ls and stat)
RUN apt-get update && apt-get install -y \
    coreutils

# Set the working directory to /app
WORKDIR /app

# Copy the inode script into the container
COPY ./inodes.sh .
# Ensure the script is executable
RUN chmod +x inodes.sh
