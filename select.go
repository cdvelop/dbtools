package dbtools

import (
	"context"
	"fmt"
)

func SelectOne(sql string, o dboAdapter, ctx context.Context) (map[string]string, error) {
	rows, err := o.Open().QueryContext(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("error QueryContext in SelectOne %v", err)
	}

	return FetchOne(rows)
}

func SelectAll(sql string, o dboAdapter, ctx context.Context) ([]map[string]string, error) {
	rows, err := o.Open().QueryContext(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("error QueryContext in SelectAll %v", err)
	}

	return FetchAll(rows)
}
