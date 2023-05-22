package dbtools

import (
	"fmt"
	"strings"

	"github.com/cdvelop/model"
)

func createSqlListByField(table model.Object) (sqlist, keyList []string) {
	var (
		defaulType     string   //TEXT tipo por defecto en base de datos
		foreignKeyList []string //FOREIGN KEY si los hay
	)

	for _, field := range table.Fields {

		// for field.Name, valueType := range field {
		defaulType = ` TEXT`
		// if field.DataType != "" {
		// defaulType = ` ` + field.DataType
		// }
		if field.Unique {
			defaulType += ` UNIQUE`
		}

		// if table.Name == "provide" {
		// 	fmt.Printf("FIELD : [%v]\n", field)
		// }

		if primaryKey, primaryKeyThisTable := IdpkTABLA(field.Name, table.Name); primaryKey {
			if primaryKeyThisTable {
				defaulType = defaulType + ` PRIMARY KEY NOT NULL`
			} else {
				defaulType = defaulType + ` NOT NULL`

				fkName := field.Name[2:] //nombre tabla foránea

				foreignTableName := strings.ReplaceAll(fkName, "_", "") //remover _
				cf := fmt.Sprintf("CONSTRAINT fk_%v FOREIGN KEY (%v) REFERENCES %v(%v) ON DELETE CASCADE",
					foreignTableName, field.Name, foreignTableName, field.Name)
				// CONSTRAINT fk_departments FOREIGN KEY (department_id) REFERENCES departments(department_id)
				foreignKeyList = append(foreignKeyList, cf)
			}
		}

		sqlist = append(sqlist, field.Name+defaulType)
		keyList = append(keyList, field.Name)
		// }
	}

	//hay FOREIGN KEY ?
	if len(foreignKeyList) > 0 {
		c := `` //coma si es necesario
		if len(foreignKeyList) > 1 {
			c = `, `
		}
		f := strings.Join(foreignKeyList, c)
		sqlist = append(sqlist, f)
	}
	return
}
