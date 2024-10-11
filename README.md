# CSV to JSON Lines Converter

## Overview

This Go command-line application converts comma-separated values (CSV) files into JSON Lines (JSONL) format. It is designed to facilitate data handling for database population and document storage.

## Features

- Converts CSV files to JSON Lines format.
- Supports command-line arguments for input and output file paths.
- Provides logging for critical operations.
- Profiling capabilities using pprof for performance monitoring.

## Requirements

- Go 1.16 or later
- Git (for version control)

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/mvellucci100/msds431Assignment2
   cd Assignment2 #This is where the go application lives

## Execution
1. Build the application
   go build -o Assignment2
2. Run the Application
   ./Assignment2 *path_to_csv* *path_to_output_file.jl*
3. Check your output file and view your json file
