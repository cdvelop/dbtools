package dbtools

import (
	"database/sql"
)

func FetchOne(rows *sql.Rows) (out map[string]string, err string) {
	const this = "FetchOne error "
	if !rows.Next() {
		return nil, ""
	}

	columns, e := rows.Columns()
	if e != nil {
		return nil, this + e.Error()
	}

	// fmt.Println("COLUMNAS: ", columns)

	columnCount := len(columns)
	row, err := ScanOne(rows, columnCount, columns)

	if err != "" {
		return nil, this + e.Error()
	}

	defer rows.Close()
	if e := rows.Close(); e != nil {
		return nil, this + e.Error()
	}
	return row, ""
}

// FetchAll .
func FetchAll(rows *sql.Rows) (out []map[string]string, err string) {
	const this = "FetchAll error"
	var columns []string
	var columnCount int
	var e error

	rowArray := make([]map[string]string, 0)
	processedRows := 0

	for rows.Next() {
		// Read columns on first row only
		if processedRows == 0 {
			columns, e = rows.Columns()
			if e != nil {
				return nil, this + e.Error()
			}
			columnCount = len(columns)
		}
		row, err := ScanOne(rows, columnCount, columns)
		if err != "" {
			return nil, this + err
		}
		rowArray = append(rowArray, row)
		processedRows++
	}
	///Sin filas: devuelve cero en lugar de un mapa vacío []
	if processedRows == 0 {
		return nil, ""
	}
	return rowArray, ""
}
