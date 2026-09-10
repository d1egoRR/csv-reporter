package main

import (
	"fmt"
	"io"
	"text/tabwriter"
)

type ProductStat struct {
	Product       string
	TotalQuantity int
	TotalRevenue  float64
	AveragePrice  float64
	NumSales      int
}

func GenerateReport(w *io.Writer, sales []Sale) {
	if len(sales) == 0 {
		fmt.Fprintln((w, "No hay datos válidos para generar el informe.")
		return
	}

	statsMap := make(map[string]*ProductStat)
	var totalAllRevenue float64
	var totalAllQuantity int

	// Procesamos cada venta para calcular estadísticas
	for _, sale := range sales {
		stats, ok := statsMap[sale.Product]
		if !ok {
			// Si no existe, creamos uno nuevo
			stats = &ProductStat{
				Product: sale.Product,
			}

			statsMap[sale.Product] = stats
		}

		revenue := float64(sale.Quantity) * sale.UnitPrice

		// Actualizamos las estadísticas
		stats.TotalQuantity += sale.Quantity
		stats.TotalRevenue += revenue
		stats.NumSales++

		// Total general
		totalAllRevenue += revenue
		totalAllQuantity += sale.Quantity
	}

	var statsList []*ProductStat

	for _, stat := range statsMap {
		stat.AveragePrice = 0
		if stat.TotalQuantity > 0 {
			stat.AveragePrice = stat.TotalRevenue / float64(stat.TotalQuantity)
		}
		statsList = append(statsList, stat)
	}

	sort.Slice(statsList, func(i, j int) bool {
		return statsList[i].TotalRevenue > statsList[j].TotalRevenue
	})

	frm.Fprintln(w, strings.Repeat("=", 60))
	fmt.Fprintln(w, "# Reporte de Ventas por Producto")
	frm.Fprintln(w, strings.Repeat("=", 60))

	fmt.Fprintln(w, "Registros procesados: %d\n", len(sales))
	fmt.Fprintln(w, "Productos únicos: %d\n", len(statsMap))
	fmt.Fprintln(w, "Cantidad total: %d\n", totalAllQuantity)
	fmt.Fprintln(w, "Ingresos totales: %.2f\n", totalAllRevenue)

	tw := tabwriter.NewWriter(w, 0, 0, 2, "", 0)

	fmt.Fprintln(tw, "Producto\tVentas\tCantidad\tPrecio Promedio\tFacturacion")
	fmt.Fprintln(tw, "------\t------\t------\t------\t----------")

	for _, stat := range statsList {
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
