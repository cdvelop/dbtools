package dbtools

import (
	"context"
)

func SelectOne(sql string, o dboAdapter, ctx context.Context) (out map[string]string, err string) {
	rows, e := o.Open().QueryContext(ctx, sql)
	if e != nil {
		return nil, "error QueryContext in SelectOne " + e.Error()
	}

	return FetchOne(rows)
}

func SelectAll(sql string, o dboAdapter, ctx context.Context) (out []map[string]string, err string) {
	rows, e := o.Open().QueryContext(ctx, sql)
	if e != nil {
		return nil, "error QueryContext in SelectAll " + e.Error()
	}

	return FetchAll(rows)
}
