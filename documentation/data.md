# Data Package

This document describes the `data` package in Golang, which provides functionalities for managing and manipulating data stored in text files. It defines an interface for data handling operations and a concrete implementation.

## Functionality

The `data` package offers the following features:

- Duplicate Removal: Eliminates duplicate lines from a data file.
- Cache Management: Creates and removes cache files associated with the data file.
## Interface: DataHandler

The `DataHandler` interface defines the operations that can be performed on a data handler object. Any concrete implementation must adhere to this interface.

Methods:

`GetFilename() string`: Returns the name of the sorted data file.
`RemoveCache() error`: Deletes the existing cache files.
`RemoveDuplicates() error`: Removes duplicate lines from the data file and stores the unique lines in a sorted cache file.
`CreateCache() error`: Creates empty cache files for the data.
## Implementation: Data struct

The `Data` struct implements the `DataHandler` interface. It stores the filenames for both the original data file and the sorted cache file.

Methods:

`GetFilename() string`: Returns the sorted data filename (`filename_sort`).
`RemoveCache() error`: Removes both the data and sorted cache files, then recreates them empty.
`CreateCache() error`: Creates empty cache files for the data and sorted data. (This implementation can be improved to handle potential errors during file creation.)
`RemoveDuplicates() error`: Opens the data file for reading, the sorted cache file for appending, and uses a map to store encountered lines. It iterates through the data file line by line, checks for duplicates in the map, and writes unique lines to the sorted cache file.

## Usage Example

    //main.go
    package main

    import (
    "fmt"

    "./data" // Import the data package
    )

    func main() {
    dataHandler := data.NewData()

    // Perform operations on the data handler
    err := dataHandler.RemoveDuplicates()
    if err != nil {
        fmt.Println("Error removing duplicates:", err)
        return
    }

    fmt.Println("Duplicates removed successfully. Sorted data is in", dataHandler.GetFilename())
    }
