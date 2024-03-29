package dbtools

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/cdvelop/model"
)

// UpdateTABLES revisa si tienen data las tablas para actualizarlas y respaldar la data
func UpdateTablesTransaction(o dboAdapter, tx *sql.Tx, ctx context.Context, tables ...*model.Object) (err string) {
	const this = "UpdateTablesTransaction err "
	for _, table := range tables {
		//consulta entrega columna nombre
		q := fmt.Sprintf(o.SQLTableInfo(), table.ObjectName)
		// tableInfo, ok := objectdb.QueryOne(q)
		tableInfo, err := TxSelectOne(q, ctx, tx)
		if err != "" {
			return this + err
		}

		if len(tableInfo) == 0 { //si no existe crear tabla nueva
			return CreateTableInTransaction(table, tx, ctx)
		} else { //revisar tabla consultar si tiene data
			list, err := TxSelectOne("SELECT * FROM "+table.ObjectName+";", ctx, tx)
			if err != "" {
				return this + err
			}

			if len(list) == 0 { //lista sin data borramos tabla y la creamos nuevamente para no chequearla
				err := DeleteTableInTransaction(table, o, tx, ctx)
				if err != "" {
					return this + err
				}

				return CreateTableInTransaction(table, tx, ctx)

			} else { //lista con data hay que actualizar
				// fmt.Printf("CLon Tabla: %v list: %v\n", table.Name, list)
				// log.Printf("tabla %v con data. hay que verificar", table.Name)
				//clonamos la tabla con data a la nueva
				err := ClonOneTableInTransaction(o, table, tx, ctx)
				if err != "" {
					return this + err
				}
			}
		}
	} //* ****tablas*****

	fmt.Println(">>> actualización de tablas ok")
	return ""
}
