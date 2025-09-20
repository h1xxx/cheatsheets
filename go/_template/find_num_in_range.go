package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"sync"
)

// Range represents a range of integers and its associated value
type Range struct {
	Start int
	End   int
	Value string
}

func main() {
	// Open the CSV file
	file, err := os.Open("find_num_in_range.data")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Skip the header
	_, err = reader.Read()
	if err != nil {
		fmt.Println("Error reading header:", err)
		return
	}

	// Parse the CSV data into a slice of Range structs
	var ranges []Range
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading record:", err)
			return
		}

		start, err := strconv.Atoi(record[0])
		if err != nil {
			fmt.Println("Error parsing range_start:", err)
			return
		}

		end, err := strconv.Atoi(record[1])
		if err != nil {
			fmt.Println("Error parsing range_end:", err)
			return
		}

		value := record[2]

		ranges = append(ranges, Range{Start: start, End: end, Value: value})
	}

	// Sort the ranges by Start
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})

	// Example: 1 million integers to look up
	inputs := []int{-10000, 42, 1001, -5, 7500, 999, 500, 2000, 10000, 300, 7000} // Replace with your 1 million integers

	// Use a wait group to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(len(inputs))

	// Process each input in parallel
	for _, input := range inputs {
		go func(input int) {
			defer wg.Done()
			value := findValueForInput(ranges, input)
			fmt.Printf("The value for %d is: %s\n", input, value)
		}(input)
	}

	// Wait for all goroutines to finish
	wg.Wait()
}

// findValueForInput performs a binary search to find the value for the given input
func findValueForInput(ranges []Range, input int) string {
	low := 0
	high := len(ranges) - 1

	for low <= high {
		mid := (low + high) / 2
		if input >= ranges[mid].Start && input <= ranges[mid].End {
			return ranges[mid].Value
		} else if input < ranges[mid].Start {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return "not_found"
}

