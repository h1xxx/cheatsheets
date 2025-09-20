package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"strings"
)

var (
	// the 'P' versions allow specifying both a long name and a short name
	name    = pflag.StringP("name", "n", "World", "Name to greet")
	count   = pflag.IntP("count", "c", 1, "Number of times to greet")
	verbose = pflag.BoolP("verbose", "v", false, "Enable verbose output")
	tags    = pflag.StringSliceP("tag", "t", []string{},
		"Tags to apply (can be specified multiple times)")
	output = pflag.StringP("output", "o", "",
		"Output file (e.g., --output=file.txt or -o file.txt)")
)

func main() {
	pflag.Parse()

	if *verbose {
		fmt.Println("Verbose mode enabled.")
		// non-flag arguments
		fmt.Printf("Raw arguments provided: %v\n", pflag.Args())
		fmt.Println("---")
	}

	for i := 0; i < *count; i++ {
		greeting := fmt.Sprintf("Hello, %s!", *name)
		if len(*tags) > 0 {
			greeting += fmt.Sprintf(" [Tags: %s]", strings.Join(*tags, ", "))
		}
		fmt.Println(greeting)
	}

	if *output != "" {
		fmt.Printf("Output would be written to: %s\n", *output)
	}

	// positional arguments
	remainingArgs := pflag.Args()
	if len(remainingArgs) > 0 {
		fmt.Println("---")
		fmt.Printf("Positional arguments: %v\n", remainingArgs)
	}
}
