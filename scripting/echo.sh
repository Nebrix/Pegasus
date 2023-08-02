#!/bin/bash

# Function definition for echoln
function echoln {
  echo "$1"
}

# Check if a file name is provided as an argument
if [ $# -eq 0 ]; then
  echo "Usage: $0 <script.ps>"
  exit 1
fi

# Execute the script using the 'source' command
source "$1"
