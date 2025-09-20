package main

import (
	"flag"
	"fmt"
)

func main() {
	// Define flags
	name := flag.String("name", "World", "a name to say hello to")
	age := flag.Int("age", 0, "your age")
	verbose := flag.Bool("v", false, "enable verbose output")

	// Parse flags
	flag.Parse()

	// Use flags
	if *verbose {
		fmt.Println("Verbose mode enabled")
	}
	fmt.Printf("Hello, %s! You are %d years old.\n", *name, *age)

	// Handle additional non-flag arguments
	args := flag.Args()
	if len(args) > 0 {
		fmt.Println("Additional args:", args)
	}
}

