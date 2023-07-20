package dbtools

import (
	"database/sql"
)

func FetchOne(rows *sql.Rows) (map[string]string, error) {
	if !rows.Next() {
		return nil, nil
	}

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	// fmt.Println("COLUMNAS: ", columns)

	columnCount := len(columns)
	row, err := ScanOne(rows, columnCount, columns)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	if err := rows.Close(); err != nil {
		return nil, err
	}
	return row, nil
}

// FetchAll .
func FetchAll(rows *sql.Rows) ([]map[string]string, error) {
	var columns []string
	var columnCount int
	var err error

	rowArray := make([]map[string]string, 0)
	processedRows := 0

	for rows.Next() {
		// Read columns on first row only
		if processedRows == 0 {
			columns, err = rows.Columns()
			if err != nil {
				return nil, err
			}
			columnCount = len(columns)
		}
		row, err := ScanOne(rows, columnCount, columns)
		if err != nil {
			return nil, err
		}
		rowArray = append(rowArray, row)
		processedRows++
	}
	///Sin filas: devuelve cero en lugar de un mapa vacío []
	if processedRows == 0 {
		return nil, nil
	}
	return rowArray, nil
}
