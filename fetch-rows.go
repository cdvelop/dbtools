package dbtools

import "database/sql"

//FetchOne .
func FetchOne(rows *sql.Rows) (map[string]string, error) {
	if !rows.Next() {
		return nil, nil
	}

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	columnCount := len(columns)
	row, err := ScanOne(rows, columnCount, columns)

	if err != nil {
		return nil, err
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}
	return row, nil
}

func FetchAll(rows *sql.Rows) ([]map[string]string, error) {
	rowArray := make([]map[string]string, 0)

	for {
		row, err := FetchOne(rows)
		if err != nil {
			return nil, err
		}
		if row == nil {
			break
		}
		rowArray = append(rowArray, row)
	}

	return rowArray, nil
}
