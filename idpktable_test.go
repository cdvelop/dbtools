package dbtools_test

import (
	"testing"

	"github.com/cdvelop/dbtools"
)

func Test_IdpkTABLA(t *testing.T) {
	// table_name := "user"

	TestData := map[string]struct {
		table_name          string
		keyNameIN           string
		primarykey          bool
		primaryKeyThisTable bool
	}{
		"id corresponde a tabla sin guion":                               {table_name: "usuario", keyNameIN: "idusuario", primarykey: true, primaryKeyThisTable: true},
		"id corresponde a tabla guion bajo _":                            {table_name: "especialidad", keyNameIN: "id_especialidad", primarykey: true, primaryKeyThisTable: true},
		"corresponde a tabla y key contiene parte el nombre de la tabla": {table_name: "especialidad", keyNameIN: "especialidades", primarykey: false, primaryKeyThisTable: false},
		"id fk de otra tabla sin guion":                                  {table_name: "usuario", keyNameIN: "idfactura", primarykey: true, primaryKeyThisTable: false},
		"id fk de otra tabla con guion bajo _":                           {table_name: "usuario", keyNameIN: "id_factura", primarykey: true, primaryKeyThisTable: false},
		"no primary key presente":                                        {table_name: "usuario", keyNameIN: "factura", primarykey: false, primaryKeyThisTable: false},
		"menos de 2 caracteres id no presente":                           {table_name: "usuario", keyNameIN: "i", primarykey: false, primaryKeyThisTable: false},
	}

	for testName, dt := range TestData {

		t.Run((testName), func(t *testing.T) {
			pk, pkthistable := dbtools.IdpkTABLA(dt.keyNameIN, dt.table_name)

			if pk != dt.primarykey {
				t.Fail()
			}
			if pkthistable != dt.primaryKeyThisTable {
				t.Fail()
			}
		})

	}
}
