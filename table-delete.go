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
func DeleteTABLE(o dboAdapter, table_name string) {
	db := o.Open()
	defer db.Close()

	sql := fmt.Sprintf(o.DropTable(), table_name)
	if _, err := db.Exec(sql); err != nil {
		log.Fatal("Error DeleteTABLE: ", err)
	}
	fmt.Printf(">>> Tabla %v eliminada\n", table_name)
}

func DeleteTableInTransaction(table *model.Object, o OrmAdapter, tx *sql.Tx, ctx context.Context) error {
	sql := fmt.Sprintf(o.SQLDropTable(), table.Name)

	_, err := tx.ExecContext(ctx, sql)
	if err != nil {
		tx.Rollback()
		return err
	}
	fmt.Printf(">>> Tabla %v eliminada\n", table.Name)
	return nil
}
