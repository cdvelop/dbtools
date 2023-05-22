package dbtools

import (
	"context"
	"database/sql"
	"log"
)

// SelectAll...
func SelectAll(sql string, st *sql.DB, ctx context.Context) (rowsMap []map[string]string, ok bool) {
	rowsMap = make([]map[string]string, 0) //inicializamos la salida

	rows, err := st.QueryContext(ctx, sql)
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
