package dbtools

import "database/sql"

type ormAdapter interface {
	SQLTableInfo() string
	SQLColumName() string
	SQLDropTable() string
}

type operation struct {
	*sql.DB
	ormAdapter
}

// ormAdapter:
// SQLTableInfo() string //sql como obtiene la base de datos el nombre de la tabla
// SQLColumName() string //sql como se llama a la columna en el motor de base de datos
// SQLDropTable() string //sql de eliminación de tabla
func NewOperationDB(db *sql.DB, orm ormAdapter) *operation {
	return &operation{
		DB:         db,
		ormAdapter: orm,
	}
}
