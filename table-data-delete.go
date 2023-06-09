package dbtools

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/cdvelop/model"
)

// DeleteDataFromTABLE borra data de una tabla en db
func DeleteDataFromTABLE(dba dbAdapter, table_name string) {
	db := dba.Open()
	defer db.Close()
	sql := fmt.Sprintf("DELETE FROM %v;", table_name)
	if _, err := db.Exec(sql); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("<<< data tabla: %v eliminada ...\n", table_name)
}

// DeleteTABLE elimina tabla de una base de datos
func DeleteTABLE(dba dbAdapter, table_name string) {
	db := dba.Open()
	defer db.Close()
	sql := fmt.Sprintf("DROP TABLE IF EXISTS %v CASCADE;", table_name)
	if _, err := db.Exec(sql); err != nil {
		log.Fatal(err)
	}
	fmt.Printf(">>> Tabla %v eliminada\n", table_name)
}

func DeleteTableInTransaction(table model.Object, o OrmAdapter, tx *sql.Tx, ctx context.Context) bool {
	sql := fmt.Sprintf(o.SQLDropTable(), table.Name)

	_, err := tx.ExecContext(ctx, sql)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
		return false
	}
	fmt.Printf(">>> Tabla %v eliminada\n", table.Name)
	return true
}
