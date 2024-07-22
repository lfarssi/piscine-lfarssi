#!/bin/bash

# Count the number of regular files
file_count=$(find . -type f | wc -l)

# Count the number of directories
dir_count=$(find . -type d | wc -l)

# Print the total count
echo $((file_count + dir_count))