package dbtools

import (
	"database/sql"
	"fmt"
)

// ScanOne
func ScanOne(rows *sql.Rows, columnCount int, columns []string) (map[string]string, error) {
	scanFrom := make([]interface{}, columnCount)
	values := make([]interface{}, columnCount)
	for i := range scanFrom {
		scanFrom[i] = &values[i]
	}
	err := rows.Scan(scanFrom...)
	if err != nil {
		return nil, err
	}
	row := make(map[string]string)
	//Construye el mapa asociativo a partir de valores y nombres de columna
	for i := range values {
		if values[i] == nil {
			row[columns[i]] = ""
			// log.Printf("valor nulo")
		} else {
			row[columns[i]] = fmt.Sprint((values[i]))
		}
		// row[columns[i]] = values[i]
	}
	return row, nil
}
