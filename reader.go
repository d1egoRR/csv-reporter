package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

type Sale struct {
	Date     time.Time
	Product  string
	Quantity int
	Price    float64
	Line     int
}

func ReadCSV(path string) ([]Sale, []string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, fmt.Errorf("Abriendo fichero: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.TrimLeadingSpace = true

	header, err := reader.Read()
	if err != nil {
		return nil, nil, fmt.Errorf("Error en cabecera: %w", err)
	}

	if len(header) < 4 {
		return nil, nil, fmt.Errorf("Cabecera incompleta")
	}

	var sales []Sale
	var warnings []string
	var lineNum int = 1

	for {
		lineNum++

		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			warnings = append(
				warnings,
				fmt.Sprintf("Línea %d: error de formato CSV: %v", lineNum, err),
			)
			continue
		}

		sale, err := parseRecord(record, lineNum)
		if err != nil {
			warnings = append(
				warnings,
				fmt.Sprintf("línea %d: %v", lineNum, err),
			)
			continue
		}

		sales = append(sales, sale)
	}

	return sales, warnings, nil
}

func parseRecord(record []string, line int) (Sale, error) {
	if len(record) < 4 {
		return Sale{}, fmt.Errorf("línea incompleta")
	}

	date, err := time.Parse("2006-01-02", record[0])
	if err != nil {
		return Sale{}, fmt.Errorf("fecha inválida")
	}

	product := strings.TrimSpace(record[1])
	if product == "" {
		return Sale{}, fmt.Errorf("producto inválido")
	}

	quantity, err := strconv.Atoi(strings.TrimSpace(record[2]))
	if err != nil {
		return Sale{}, fmt.Errorf("cantidad inválida")
	}

	if quantity <= 0 {
		return Sale{}, fmt.Errorf("cantidad debe ser mayor a cero")
	}

	price, err := strconv.ParseFloat(strings.TrimSpace(record[3]), 64)
	if err != nil {
		return Sale{}, fmt.Errorf("precio inválido")
	}

	if price < 0 {
		return Sale{}, fmt.Errorf("precio negativo")
	}

	sale := Sale{
		Date:     date,
		Product:  product,
		Quantity: quantity,
		Price:    price,
		Line:     line,
	}

	return sale, nil
}
