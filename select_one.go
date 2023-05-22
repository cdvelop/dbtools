package dbtools

import (
	"context"
	"database/sql"
	"log"
)

func SelectOne(sql string, st *sql.DB, ctx context.Context) (rowsMap map[string]string, ok bool) {
	rowsMap = make(map[string]string, 0) //inicializamos la salida

	rows, err := st.QueryContext(ctx, sql)
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
