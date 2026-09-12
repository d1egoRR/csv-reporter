# CSV Reporter

Aplicación CLI en Go para procesar archivos CSV de ventas, validar registros y generar reportes de facturación tabulados.

## Requisitos

- [Go](https://go.dev/) 1.20 o superior

## Compilación y Ejecución

```bash
# Compilar ejecutable
go build -o csv-reporter

# O ejecutar directamente el paquete
go run . --input sample.csv
```

> **Nota:** Se utiliza `go run .` (y no `go run main.go`) para compilar todos los archivos del paquete `main`.

## Flags disponibles

| Flag | Descripción | Por defecto | Requerido |
| --- | --- | --- | --- |
| `--input` | Ruta del archivo CSV a procesar | `""` | **Sí** |
| `--output` | Ruta del archivo de reporte | `"report.txt"` | No |
| `--verbose` | Muestra detalles del proceso en consola | `false` | No |
| `--version` | Muestra la versión de la aplicación | `false` | No |

## Ejemplos de uso

### Comandos Cobra

```bash
# Generar informe de ventas con Cobra
go run . report --input sample.csv --output informe_cobra.txt

# Validar archivo CSV sin generar informe
go run . validate --input sample.csv
```

### Modo tradicional (flags)

```bash
# Procesar archivo básico
go run . --input sample.csv

# Especificar archivo de salida y modo verbose
go run . --input sample.csv --output reporte.txt --verbose
```

## Estructura del Proyecto

- `main.go` - Entrada principal y ejecución de comandos.
- `cobra_command_sample.go` - Definición de comandos y flags con Cobra (`report`, `validate`).
- `reader.go` - Lectura y parseo de registros CSV (`fecha,producto,cantidad,precio`).
- `validator.go` - Validación de datos (fechas futuras, límites de precio y cantidad).
- `reporter.go` - Generación de estadísticas y formato de reporte tabulado (`ReportSummary`).