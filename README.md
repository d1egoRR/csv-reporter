# csv-reporter

`csv-reporter` is a command-line interface (CLI) application written in Go designed to process CSV files and output report files.

This project follows the tutorial guide from [Oshy.tech - Go CLI Tutorial](https://oshy.tech/blog/cli-go/).

## Features

- **CLI Flag Parsing**: Uses standard Go `flag` package for intuitive command-line flag handling.
- **Input Validation**: Ensures mandatory parameters are supplied before processing.
- **Customizable Output**: Specify custom output filenames with a default fallback (`report.txt`).
- **Verbose Output**: Optional verbose mode for detailed execution logging.

## Prerequisites

- [Go](https://go.dev/doc/install) (version 1.20 or later recommended)

## Installation & Build

Clone the repository and build the binary:

```bash
git clone https://github.com/d1egoRR/csv-reporter.git
cd csv-reporter
go build -o csv-reporter
```

Alternatively, you can run the application directly using `go run`.

## Usage

### Command Line Flags

| Flag | Description | Default | Required |
| --- | --- | --- | --- |
| `-input` | Path to the input CSV file | `""` | **Yes** |
| `-output` | Path to the generated output file | `"report.txt"` | No |
| `-verbose` | Enable verbose output logging | `false` | No |

### Usage Examples

#### Display Help & Options
```bash
go run main.go -help
```

#### Process a CSV file (Default output: `report.txt`)
```bash
go run main.go -input data.csv
```

#### Process a CSV file with custom output file
```bash
go run main.go -input data.csv -output output_report.txt
```

#### Process a CSV file with Verbose Mode enabled
```bash
go run main.go -input data.csv -output output_report.txt -verbose
```

### Example Execution Output

```bash
$ go run main.go -input data.csv -output result.txt -verbose
Input file: data.csv
Output file: result.txt
Processing file: data.csv -> result.txt
```

## Project Structure

- `main.go` - Main entry point and CLI flag parsing logic.
- `go.mod` - Go module declaration (`github.com/d1egoRR/csv-reporter`).
- `README.md` - Project documentation.

## License

This project is open source and available under the [MIT License](LICENSE).