package dbtools

import (
	"fmt"
	"log"

	"github.com/cdvelop/model"
)

// UpdateTABLES revisa si tienen data las tablas para actualizarlas y respaldar la data
func UpdateTABLES(dba dbAdapter, o OrmAdapter, tables ...model.Object) bool {
	db := dba.Open()
	defer db.Close()

	for _, table := range tables {

		//consulta entrega columna nombre
		q := fmt.Sprintf(o.SQLTableInfo(), table.Name)

		rows, err := db.Query(q)
		if err != nil {
			fmt.Println(err)
			return false
		}

		tableInfo, err := FetchOne(rows)
		if err != nil {
			fmt.Println(err)
			return false
		}

		if len(tableInfo) == 0 { //si no existe crear tabla nueva
			CreateOneTABLE(dba, table)
		} else { //revisar tabla consultar si tiene data

			rows, err := db.Query("SELECT * FROM " + table.Name + ";")
			if err != nil {
				fmt.Println(err)
				return false
			}

			list, err := FetchOne(rows)
			if err != nil {
				fmt.Println(err)
				return false
			}

			if len(list) == 0 { //lista sin data borramos tabla y la creamos nuevamente para no chequearla
				q := fmt.Sprintf(o.SQLDropTable(), table.Name)

				fmt.Printf(">>> Borrando tabla: %v", table.Name)

				if _, err := db.Exec(q); err != nil {
					log.Fatalf("!!! Error al borrar tabla DROP TABLE: %v %v", table.Name, err)
					return false
				}

				fmt.Printf(">>> tabla %v sin data borrada\n", table.Name)

				if !CreateOneTABLE(dba, table) {
					return false
				}
				fmt.Printf(">>> tabla %v creada\n", table.Name)

			} else { //lista con data hay que actualizar
				// fmt.Printf("CLon Tabla: %v list: %v\n", table.Name, list)
				// log.Printf("tabla %v con data. hay que verificar", table.Name)
				if !ClonDATABLE(dba, o, table) { //clonamos la tabla con data a la nueva
					log.Fatalf("!!! error al copiar la data tabla " + table.Name)
					return false
				}
			}

		}

	} //* ****tablas*****

	fmt.Println(">>> actualización de tablas completada")
	return true
}
