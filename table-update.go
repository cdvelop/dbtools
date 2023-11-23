package dbtools

import (
	"fmt"

	"github.com/cdvelop/model"
)

// UpdateTABLES revisa si tienen data las tablas para actualizarlas y respaldar la data
func UpdateTABLES(o dboAdapter, tables ...*model.Object) (err string) {
	const this = "UpdateTABLES error"
	db := o.Open()
	defer db.Close()

	for _, table := range tables {

		//consulta entrega columna nombre
		q := fmt.Sprintf(o.SQLTableInfo(), table.ObjectName)

		rows, e := db.Query(q)
		if e != nil {
			return this + e.Error()
		}

		tableInfo, err := FetchOne(rows)
		if err != "" {
			return err
		}

		if len(tableInfo) == 0 { //si no existe crear tabla nueva
			err = CreateOneTABLE(o, table)
		} else { //revisar tabla consultar si tiene data

			rows, e := db.Query("SELECT * FROM " + table.ObjectName + ";")
			if e != nil {
				return this + e.Error()
			}

			list, err := FetchOne(rows)
			if err != "" {
				return this + err
			}

			if len(list) == 0 { //lista sin data borramos tabla y la creamos nuevamente para no chequearla
				q := fmt.Sprintf(o.SQLDropTable(), table.ObjectName)

				fmt.Printf(">>> Borrando tabla: %v", table.ObjectName)

				if _, e := db.Exec(q); e != nil {
					return "!!! " + this + " al borrar tabla DROP TABLE: " + table.ObjectName + " " + e.Error()
				}

				fmt.Printf(">>> tabla %v sin data borrada\n", table.ObjectName)

				err := CreateOneTABLE(o, table)
				if err != "" {
					return this + err
				}

				fmt.Printf(">>> tabla %v creada\n", table.ObjectName)

			} else { //lista con data hay que actualizar
				// fmt.Printf("CLon Tabla: %v list: %v\n", table.Name, list)
				// log.Printf("tabla %v con data. hay que verificar", table.Name)
				//clonamos la tabla con data a la nueva
				err := ClonDATABLE(o, table)
				if err != "" {
					return this + err
				}
			}
		}
	} //* ****tablas*****

	fmt.Println(">>> actualización de tablas completada")
	return
}
