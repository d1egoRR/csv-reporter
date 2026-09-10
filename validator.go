package main

import (
	"fmt"
	"time"
)

type ValidationResult struct {
	Valid   []Sale
	Invalid []string
}

func ValidateSales(sales []Sale) ValidationResult {
	var result ValidationResult
	now := time.Now()

	for _, s := range sales {
		if s.Date.After(now) {
			result.Invalid = append(
				result.Invalid,
				fmt.Sprintf("Linea %d: La fecha no puede ser futura", s.Line),
			)
			continue
		}

		if s.Price > 10000 {
			result.Invalid = append(
				result.Invalid,
				fmt.Sprintf("Linea %d: El precio no puede ser mayor a 10000", s.Line),
			)
			continue
		}

		if s.Quantity > 1000 {
			result.Invalid = append(
				result.Invalid,
				fmt.Sprintf("Linea %d: La cantidad no puede ser mayor a 1000", s.Line),
			)
			continue
		}

		result.Valid = append(result.Valid, s)
	}

	return result
}
