#!/bin/bash

# Compile the Go program
go build */*.go

# Check if compilation was successful
if [ $? -eq 0 ]; then
    echo "Compilation successful"

    # Run tests for the Go program
    go test ./...

    # Check if tests passed
    if [ $? -eq 0 ]; then
        echo "All tests passed, running the program"
        
        # Run the compiled Go program
        ./main
    else
        echo "Tests failed, please address the issues before running the program"
    fi
else
    echo "Compilation failed"
fi
