package dbtools

import (
	"context"
	"log"
)

func SelectOne(sql string, dba dbAdapter, ctx context.Context) (rowsMap map[string]string, ok bool) {
	db := dba.Open()
	defer db.Close()

	rowsMap = make(map[string]string, 0) //inicializamos la salida

	rows, err := db.QueryContext(ctx, sql)
	if err != nil {
		log.Println(err)
		return
	}

	rowsMap, err = FetchOne(rows)
	if err != nil {
		log.Println(err)
		return
	}

	ok = true
	return
}
