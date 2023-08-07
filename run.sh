#!/bin/bash

# Compile the Go program
go build

# Check if compilation was successful
if [ $? -eq 0 ]; then
    echo "Compilation successful"

    # Run tests for the Go program
    go test ./...

    # Check if tests passed
    if [ $? -eq 0 ]; then
        echo "All tests passed! Running..."
        
        # Run the compiled Go program
        ./webgo
    else
        echo "Tests failed, please address the issues before running"
    fi
else
    echo "Compilation failed"
fi
