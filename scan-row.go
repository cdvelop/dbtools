package dbtools

import (
	"database/sql"
	"fmt"
)

func ScanOne(rows *sql.Rows, columnCount int, columns []string) (row map[string]string, err string) {
	const this = "ScanOne error "
	scanFrom := make([]interface{}, columnCount)
	values := make([]interface{}, columnCount)
	for i := range scanFrom {
		scanFrom[i] = &values[i]
	}
	e := rows.Scan(scanFrom...)
	if e != nil {
		return nil, this + e.Error()
	}

	// fmt.Println("VALUES: ", values)

	row = make(map[string]string)
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
	return row, ""
}
