package dbtools

import (
	"context"
	"log"
)

// SelectAll...
func SelectAll(sql string, dba dbAdapter, ctx context.Context) (rowsMap []map[string]string, ok bool) {
	db := dba.Open()
	defer db.Close()

	rowsMap = make([]map[string]string, 0) //inicializamos la salida

	rows, err := db.QueryContext(ctx, sql)
	if err != nil {
		log.Println("ERROR QueryContext " + err.Error())
		return
	}

	rowsMap, err = FetchAll(rows)
	if err != nil {
		log.Println("ERROR FetchAll " + err.Error())
		return
	}

	ok = true
	return
}
