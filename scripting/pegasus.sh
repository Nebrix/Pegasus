#!/bin/bash

# Function definition for echoln
function echoln {
  local message="$1"
  local pattern='[$][A-Za-z_][A-Za-z_0-9]*'
  while [[ $message =~ $pattern ]]; do
    local var="${BASH_REMATCH[0]}"
    local var_name="${var:1}"
    local var_value="${!var_name}"
    message="${message//$var/$var_value}"
  done
  echo "$message"
}

function readln {
  read -p "$1" value
  #echo "$value"
}

function let {
  local var_name="$1"
  local var_value="$2"
  # Use eval to perform variable assignment
  eval "$var_name=\"$var_value\""
}

# Check if a file name is provided as an argument
if [ $# -eq 0 ]; then
  echo "Usage: $0 <script.pg>"
  exit 1
fi

# Execute the script using the 'source' command
source "$1"
