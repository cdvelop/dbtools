package dbtools

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

func RenameTable(old_name, new_name string, tx *sql.Tx, ctx context.Context) bool {

	_, err := tx.ExecContext(ctx, fmt.Sprintf("ALTER TABLE %v RENAME TO %v;", old_name, new_name))
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return false
	}

	fmt.Printf(">>> Nombre Tabla: [%v] cambiada a: [%v]\n", old_name, new_name)
	return true
}

func RenameColumn(table_name, old_column, new_column string, tx *sql.Tx, ctx context.Context) bool {

	_, err := tx.ExecContext(ctx, fmt.Sprintf("ALTER TABLE %v RENAME COLUMN %v TO %v;", table_name, old_column, new_column))
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return false
	}

	fmt.Printf(">>> Columna: [%v] de la Tabla: [%v] se cambio a: [%v]\n", old_column, table_name, new_column)
	return true
}

func AddColumn(table_name, column_name, data_type string, tx *sql.Tx, ctx context.Context) bool {

	_, err := tx.ExecContext(ctx, fmt.Sprintf("ALTER TABLE %v ADD COLUMN %v %v;", table_name, column_name, data_type))
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return false
	}

	fmt.Printf(">>> Se agrego la Columna: [%v] tipo: [%v] a la Tabla: [%v]\n", column_name, data_type, table_name)
	return true

}
