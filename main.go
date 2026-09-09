package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	input := flag.String("input", "", "input csv file (required)")
	output := flag.String("output", "report.txt", "output file")
	verbose := flag.Bool("verbose", false, "show verbose output")

	flag.Parse()

	if *input == "" {
		fmt.Fprintln(os.Stderr, "Error: flag --input is required")
		flag.Usage()
		os.Exit(1)
	}

	if *verbose {
		fmt.Printf("Input file: %s\n", *input)
		fmt.Printf("Output file: %s\n", *output)
	}

	fmt.Printf("Processing file: %s -> %s\n", *input, *output)
}
