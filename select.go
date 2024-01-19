package dbtools

import (
	"context"
	"database/sql"
)

func TxSelectOne(sql string, ctx context.Context, tx *sql.Tx) (out map[string]string, err string) {
	rows, e := tx.QueryContext(ctx, sql)
	if e != nil {
		return nil, "error QueryContext in SelectOne " + e.Error()
	}

	return FetchOne(rows)
}

func TxSelectAll(sql string, ctx context.Context, tx *sql.Tx) (out []map[string]string, err string) {
	rows, e := tx.QueryContext(ctx, sql)
	if e != nil {
		return nil, "error QueryContext in SelectAll " + e.Error()
	}

	return FetchAll(rows)
}
