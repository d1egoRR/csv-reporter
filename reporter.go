package main

import (
	"cmp"
	"fmt"
	"io"
	"slices"
	"strings"
	"text/tabwriter"
)

type ProductStat struct {
	Product       string
	TotalQuantity int
	TotalRevenue  float64
	AveragePrice  float64
	NumSales      int
}

type ReportSummary struct {
	ProcessedRecords int
	UniqueProducts   int
	TotalQuantity    int
	TotalRevenue     float64
	ProductStats     []*ProductStat
}

func GenerateReport(w io.Writer, sales []Sale) {
	if len(sales) == 0 {
		fmt.Fprintln(w, "No hay datos válidos para generar el informe.")
		return
	}

	summary := processSales(sales)
	generateTableReport(w, summary)
}

func processSales(sales []Sale) (summary ReportSummary) {
	statsMap := make(map[string]*ProductStat)
	var totalAllRevenue float64
	var totalAllQuantity int

	for _, sale := range sales {
		stats, ok := statsMap[sale.Product]
		if !ok {
			// Si no existe, creamos uno nuevo
			stats = &ProductStat{
				Product: sale.Product,
			}

			statsMap[sale.Product] = stats
		}

		revenue := float64(sale.Quantity) * sale.Price

		// Actualizamos las estadísticas
		stats.TotalQuantity += sale.Quantity
		stats.TotalRevenue += revenue
		stats.NumSales++

		// Total general
		totalAllRevenue += revenue
		totalAllQuantity += sale.Quantity
	}

	summary = ReportSummary{
		ProcessedRecords: len(sales),
		UniqueProducts:   len(statsMap),
		TotalQuantity:    totalAllQuantity,
		TotalRevenue:     totalAllRevenue,
		ProductStats:     buildStatsList(statsMap),
	}

	sortStats(summary.ProductStats)

	return summary
}

func buildStatsList(statsMap map[string]*ProductStat) []*ProductStat {
	var statsList []*ProductStat

	for _, stat := range statsMap {
		stat.AveragePrice = 0
		if stat.TotalQuantity > 0 {
			stat.AveragePrice = stat.TotalRevenue / float64(stat.TotalQuantity)
		}
		statsList = append(statsList, stat)
	}

	return statsList
}

func sortStats(statsList []*ProductStat) {
	slices.SortFunc(statsList, func(a, b *ProductStat) int {
		return cmp.Compare(b.TotalRevenue, a.TotalRevenue)
	})
}

func generateTableReport(w io.Writer, summary ReportSummary) {
	fmt.Fprintln(w, strings.Repeat("=", 60))
	fmt.Fprintln(w, "# Reporte de Ventas por Producto")
	fmt.Fprintln(w, strings.Repeat("=", 60))

	fmt.Fprintf(w, "Registros procesados: %d\n", summary.ProcessedRecords)
	fmt.Fprintf(w, "Productos únicos: %d\n", summary.UniqueProducts)
	fmt.Fprintf(w, "Cantidad total: %d\n", summary.TotalQuantity)
	fmt.Fprintf(w, "Ingresos totales: %.2f\n", summary.TotalRevenue)

	tw := tabwriter.NewWriter(w, 0, 0, 2, ' ', 0)

	fmt.Fprintln(tw, "Producto\tVentas\tCantidad\tPrecio Promedio\tFacturacion")
	fmt.Fprintln(tw, "------\t------\t------\t------\t----------")

	for _, stat := range summary.ProductStats {
		fmt.Fprintf(tw, "%s\t%d\t%d\t%.2f\t%.2f\n",
			stat.Product,
			stat.NumSales,
			stat.TotalQuantity,
			stat.AveragePrice,
			stat.TotalRevenue,
		)
	}

	tw.Flush()
}
