package dbtools

import (
	"fmt"

	"github.com/cdvelop/model"
)

// UpdateTABLES revisa si tienen data las tablas para actualizarlas y respaldar la data
func UpdateTABLES(o dboAdapter, tables ...*model.Object) error {
	db := o.Open()
	defer db.Close()

	for _, table := range tables {

		//consulta entrega columna nombre
		q := fmt.Sprintf(o.SQLTableInfo(), table.Name)

		rows, err := db.Query(q)
		if err != nil {
			return err
		}

		tableInfo, err := FetchOne(rows)
		if err != nil {
			return err
		}

		if len(tableInfo) == 0 { //si no existe crear tabla nueva
			CreateOneTABLE(o, table)
		} else { //revisar tabla consultar si tiene data

			rows, err := db.Query("SELECT * FROM " + table.Name + ";")
			if err != nil {
				return err
			}

			list, err := FetchOne(rows)
			if err != nil {
				return err
			}

			if len(list) == 0 { //lista sin data borramos tabla y la creamos nuevamente para no chequearla
				q := fmt.Sprintf(o.SQLDropTable(), table.Name)

				fmt.Printf(">>> Borrando tabla: %v", table.Name)

				if _, err := db.Exec(q); err != nil {
					return fmt.Errorf("!!! Error al borrar tabla DROP TABLE: %v %v", table.Name, err)
				}

				fmt.Printf(">>> tabla %v sin data borrada\n", table.Name)

				err := CreateOneTABLE(o, table)
				if err != nil {
					return err
				}

				fmt.Printf(">>> tabla %v creada\n", table.Name)

			} else { //lista con data hay que actualizar
				// fmt.Printf("CLon Tabla: %v list: %v\n", table.Name, list)
				// log.Printf("tabla %v con data. hay que verificar", table.Name)
				//clonamos la tabla con data a la nueva
				err := ClonDATABLE(o, table)
				if err != nil {
					return err
				}
			}
		}
	} //* ****tablas*****

	fmt.Println(">>> actualización de tablas completada")
	return nil
}
