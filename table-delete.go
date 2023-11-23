package dbtools

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/cdvelop/model"
)

// DeleteDataFromTABLE borra data de una tabla en db
func DeleteDataFromTABLE(dba dbAdapter, table_name string) (err string) {
	db := dba.Open()
	defer db.Close()
	sql := fmt.Sprintf("DELETE FROM %v;", table_name)
	if _, e := db.Exec(sql); e != nil {
		return "DeleteDataFromTABLE error " + e.Error()
	}
	fmt.Printf("<<< data tabla: %v eliminada ...\n", table_name)
	return ""
}

// DeleteTABLE elimina tabla de una base de datos
func DeleteTABLE(o dboAdapter, table_name string) (err string) {
	db := o.Open()
	defer db.Close()

	sql := fmt.Sprintf(o.DropTable(), table_name)
	if _, e := db.Exec(sql); e != nil {
		return "DeleteTABLE error " + e.Error()
	}
	fmt.Printf(">>> Tabla %v eliminada\n", table_name)
	return ""
}

func DeleteTableInTransaction(table *model.Object, o OrmAdapter, tx *sql.Tx, ctx context.Context) (err string) {
	sql := fmt.Sprintf(o.SQLDropTable(), table.ObjectName)

	_, e := tx.ExecContext(ctx, sql)
	if e != nil {
		tx.Rollback()
		return "DeleteTableInTransaction error " + e.Error()
	}
	fmt.Printf(">>> Tabla %v eliminada\n", table.ObjectName)
	return ""
}
