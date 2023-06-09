package dbtools

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/cdvelop/model"
)

// UpdateTABLES revisa si tienen data las tablas para actualizarlas y respaldar la data
func UpdateTablesTransaction(dba dbAdapter, o OrmAdapter, tx *sql.Tx, ctx context.Context, tables ...model.Object) bool {

	for _, table := range tables {
		//consulta entrega columna nombre
		q := fmt.Sprintf(o.SQLTableInfo(), table.Name)
		// tableInfo, ok := objectdb.QueryOne(q)
		tableInfo, ok := SelectOne(q, dba, ctx)

		if !ok {
			return false
		}

		if len(tableInfo) == 0 { //si no existe crear tabla nueva
			if !CreateTableInTransaction(table, tx, ctx) {
				return false
			}
		} else { //revisar tabla consultar si tiene data
			if list, ok := SelectOne("SELECT * FROM "+table.Name+";", dba, ctx); ok {

				if len(list) == 0 { //lista sin data borramos tabla y la creamos nuevamente para no chequearla
					if DeleteTableInTransaction(table, o, tx, ctx) {

						if !CreateTableInTransaction(table, tx, ctx) {
							return false
						}
						fmt.Printf(">>> tabla %v creada\n", table.Name)
					} else {
						log.Fatalf("!!! error al borrar tabla DROP TABLE: " + table.Name)
						return false
					}

				} else { //lista con data hay que actualizar
					// fmt.Printf("CLon Tabla: %v list: %v\n", table.Name, list)
					// log.Printf("tabla %v con data. hay que verificar", table.Name)
					if !ClonOneTableInTransaction(dba, o, table, tx, ctx) { //clonamos la tabla con data a la nueva
						log.Fatalf("!!! error al copiar la data tabla " + table.Name)
						return false
					}

				}

			} else {
				return false
			}
		}

	} //* ****tablas*****

	fmt.Println(">>> actualización de tablas ok")
	return true
}
