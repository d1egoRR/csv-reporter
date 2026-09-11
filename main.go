package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

const version = "0.1.0"

func main() {
	var input *string = flag.String("input", "", "input csv file (required)")
	var output *string = flag.String("output", "report.txt", "output file")
	var verbose *bool = flag.Bool("verbose", false, "show verbose output")
	var showVersion *bool = flag.Bool("version", false, "show version")

	flag.Parse()

	if *showVersion {
		fmt.Println("Version:", version)
		os.Exit(0)
	}

	if *input == "" {
		fmt.Fprintln(os.Stderr, "Error: flag --input is required")
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, "Usage: go run . --input <input> [--output <output>] [--verbose] [--version]")
		os.Exit(1)
	}

	sales, warnings, err := ReadCSV(*input)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error processing file:", err)
		os.Exit(1)
	}

	if *verbose {
		fmt.Printf("Input file: %s\n", *input)
		fmt.Printf("Output file: %s\n", *output)
	}

	// show warnings
	for _, w := range warnings {
		fmt.Fprintf(os.Stderr, "[WARN] %s\n", w)
	}

	// validar
	result := ValidateSales(sales)

	if *verbose {
		fmt.Fprintf(os.Stderr, "Filas validas: %d\n", len(result.Valid))
		fmt.Fprintf(os.Stderr, "Filas invalidas: %d\n", len(result.Invalid))
	}

	for _, inv := range result.Invalid {
		fmt.Fprintf(os.Stderr, "[INV] %s\n", inv)
	}

	// generate report
	var writer io.Writer

	if *output != "" {
		f, err := os.Create(*output)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creando fichero: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()
		writer = f
	} else {
		writer = os.Stdout
	}

	GenerateReport(writer, result.Valid)

	fmt.Printf("Processing file: %s -> %s\n", *input, *output)
}
