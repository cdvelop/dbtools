package dbtools

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/cdvelop/model"
)

// CreateAllTABLES crea todas las tablas de la base de datos
func CreateAllTABLES(dba dbAdapter, tables ...model.Object) (ok bool) {
	db := dba.Open()
	defer db.Close()

	var sql []string

	if len(tables) == 0 {
		log.Fatalln("Error No hay tablas ingresadas a AddAllTablesToIndexDb")
	}

	//todo el sql por tabla
	for _, table := range tables { //array tablas ordenadas
		// fmt.Printf("Nombre de la tabla [%v]\n", table.MainName())
		sql = append(sql, makeSQLCreaTABLE(table))
	}

	q := strings.Join(sql, "\n")
	// log.Printf(">>> sql final %v", q)

	if _, err := db.Exec(q); err != nil {
		log.Fatalf("ERROR EN LA CREACIÓN DE TABLAS EN BASE DE DATOS, FUNCIÓN: CreateAllTABLES %v", err)
		return
	}

	ok = true
	return
}

// CreateOneTABLE según nombre tabla y solo con un id_nombretabla correlativo por defecto
func CreateOneTABLE(dba dbAdapter, table model.Object) bool {
	db := dba.Open()
	defer db.Close()

	sql := makeSQLCreaTABLE(table)

	if _, err := db.Exec(sql); err != nil {
		log.Fatalf("Error al Crear tabla %v %v", table.Name, err)
		return false
	}

	fmt.Printf(">>> Tabla: " + table.Name + " creada")

	return true
}

func CreateTableInTransaction(table model.Object, tx *sql.Tx, ctx context.Context) bool {
	sqlNewTable := makeSQLCreaTABLE(table)
	_, err := tx.ExecContext(ctx, sqlNewTable)
	if err != nil {
		tx.Rollback()
		return false
	}

	fmt.Printf(">>> Creando tabla: %v en db", table.Name)
	return true
}

// makeSQLCreaTABLE crea string sql crea tabla
func makeSQLCreaTABLE(table model.Object) (sql string) {
	keyLisTO, _ := createSqlListByField(table)

	column := strings.Join(keyLisTO, ", ")
	// fmt.Printf("colum tabla: %v  %v\n", table_name, column)
	return makesqlcretetable(table.Name, column)
}

// sql de creación de tabla
func makesqlcretetable(table_name, column string) (sql string) {
	sql = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %v (%v);", table_name, column)
	return
}
