package dbtools

import (
	"context"
	"fmt"
)

func SelectOne(sql string, dba dbAdapter, ctx context.Context) (map[string]string, error) {
	db := dba.Open()
	defer db.Close()

	rows, err := db.QueryContext(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("error QueryContext in SelectOne %v", err)
	}

	return FetchOne(rows)
}

// SelectAll...
func SelectAll(sql string, dba dbAdapter, ctx context.Context) ([]map[string]string, error) {
	db := dba.Open()
	defer db.Close()

	rows, err := db.QueryContext(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("error QueryContext in SelectAll %v", err)
	}

	return FetchAll(rows)
}
