#!/bin/bash

# Script to activate or create the virtual environment

ENV_DIR=".env"

# Function to print error message and exit
function error_exit {
  echo "$1" 1>&2
  exit 1
}

# Check if Python is installed
command -v python3 >/dev/null 2>&1 || error_exit "Python3 is required but it's not installed. Aborting."

# Create or activate the virtual environment
if [ -d "$ENV_DIR" ]; then
  source "$ENV_DIR/bin/activate" || error_exit "Failed to activate the virtual environment."
  echo "Virtual environment activated."
else
  echo "Virtual environment directory '$ENV_DIR' not found. Creating a new virtual environment."
  python -m venv "$ENV_DIR" || error_exit "Failed to create the virtual environment."
  source "$ENV_DIR/bin/activate" || error_exit "Failed to activate the virtual environment after creation."
  echo "Virtual environment created and activated."
fi
