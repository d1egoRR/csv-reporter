package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "csv-reporter",
	Short:   "Genera reportes a partir de archivos CSV",
	Version: "0.0.1",
}

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Genera un informe de ventas",
	RunE: func(cmd *cobra.Command, args []string) error {
		input, _ := cmd.Flags().GetString("input")
		output, _ := cmd.Flags().GetString("output")

		sales, warnings, err := ReadCSV(input)
		if err != nil {
			return err
		}

		for _, w := range warnings {
			fmt.Fprintf(os.Stderr, "[WARN] %s\n", w)
		}

		result := ValidateSales(sales)

		for _, inv := range result.Invalid {
			fmt.Fprintf(os.Stderr, "[INV] %s\n", inv)
		}

		writer := os.Stdout
		if output != "" {
			f, err := os.Create(output)
			if err != nil {
				return fmt.Errorf("Error creando file de salida %q: %v", output, err)
			}
			defer f.Close()
			writer = f
		}

		GenerateReport(writer, result.Valid)

		return nil
	},
}

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Valida un archivo CSV sin generar el informe",
	RunE: func(cmd *cobra.Command, args []string) error {
		input, _ := cmd.Flags().GetString("input")

		sales, warnings, err := ReadCSV(input)
		if err != nil {
			return err
		}

		for _, w := range warnings {
			fmt.Fprintf(os.Stderr, "[WARN] %s\n", w)
		}

		result := ValidateSales(sales)

		if len(result.Invalid) == 0 && len(warnings) == 0 {
			fmt.Println("Todas las filas son validas")
		} else {
			for _, inv := range result.Invalid {
				fmt.Println(inv)
			}
		}

		fmt.Println("")
		fmt.Println("Filas validas: ", len(result.Valid))
		fmt.Println("Filas invalidas: ", len(result.Invalid))

		return nil
	},
}

func init() {
	reportCmd.Flags().StringP("input", "i", "", "fichero CSV de entrada")
	reportCmd.Flags().StringP("output", "o", "", "fichero de salida")
	reportCmd.MarkFlagRequired("input")

	validateCmd.Flags().StringP("input", "i", "", "fichero CSV de entrada")
	validateCmd.MarkFlagRequired("input")

	rootCmd.AddCommand(reportCmd)
	rootCmd.AddCommand(validateCmd)
}
