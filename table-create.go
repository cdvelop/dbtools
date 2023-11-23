package dbtools

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/cdvelop/model"
)

// CreateTablesInDB crea todas las tablas de la base de datos
func CreateTablesInDB(dba dbAdapter, tables ...*model.Object) (err string) {
	db := dba.Open()
	defer db.Close()

	var sql []string

	if len(tables) == 0 {
		return "error no hay objetos ingresados para crear tablas"
	}

	//todo el sql por tabla
	for _, table := range tables { //array tablas ordenadas
		// fmt.Printf("Nombre de la tabla [%v]\n", table.MainName())
		sql = append(sql, makeSQLCreaTABLE(table))
	}

	q := strings.Join(sql, "\n")
	// log.Printf(">>> sql final %v", q)

	if _, e := db.Exec(q); e != nil {
		return "ERROR EN LA CREACIÓN DE TABLAS EN BASE DE DATOS, FUNCIÓN: CreateTablesInDB " + e.Error()
	}

	return ""
}

// CreateOneTABLE según nombre tabla y solo con un id_nombretabla correlativo por defecto
func CreateOneTABLE(dba dbAdapter, table *model.Object) (err string) {
	db := dba.Open()
	defer db.Close()

	sql := makeSQLCreaTABLE(table)

	if _, e := db.Exec(sql); e != nil {
		return "error al crear tabla " + table.Table + " " + e.Error()
	}

	fmt.Println(">>> Tabla: " + table.Table + " creada")

	return ""
}

func CreateTableInTransaction(table *model.Object, tx *sql.Tx, ctx context.Context) (err string) {
	sqlNewTable := makeSQLCreaTABLE(table)
	_, e := tx.ExecContext(ctx, sqlNewTable)
	if e != nil {
		tx.Rollback()
		return e.Error()
	}

	fmt.Printf(">>> Creando tabla: %v en db\n", table.Table)
	return ""
}

// makeSQLCreaTABLE crea string sql crea tabla
func makeSQLCreaTABLE(table *model.Object) string {
	keyLisTO, _ := createSqlListByField(table)

	column := strings.Join(keyLisTO, ", ")
	// fmt.Printf("colum tabla: %v  %v\n", table_name, column)
	return makesqlcretetable(table.Table, column)
}

// sql de creación de tabla
func makesqlcretetable(table_name, column string) string {
	return fmt.Sprintf("CREATE TABLE IF NOT EXISTS %v (%v);", table_name, column)
}
