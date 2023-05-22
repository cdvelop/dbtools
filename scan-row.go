package dbtools

import (
	"database/sql"
	"fmt"
)

// ScanOne
func ScanOne(rows *sql.Rows, columnCount int, columns []string) (map[string]string, error) {
	values := make([]interface{}, columnCount)
	for i := range values {
		values[i] = new(interface{})
	}
	err := rows.Scan(values...)
	if err != nil {
		return nil, err
	}
	row := make(map[string]string, columnCount)
	for i, value := range values {
		if *value.(*interface{}) == nil {
			row[columns[i]] = ""
		} else {
			row[columns[i]] = fmt.Sprintf("%v", *value.(*interface{}))
		}
	}
	return row, nil
}
