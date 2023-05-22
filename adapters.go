package dbtools

import "database/sql"

type OrmAdapter interface {
	SQLTableInfo() string
	SQLColumName() string
	SQLDropTable() string
}

type operation struct {
	*sql.DB
	OrmAdapter
}

// OrmAdapter:
// SQLTableInfo() string //sql como obtiene la base de datos el nombre de la tabla
// SQLColumName() string //sql como se llama a la columna en el motor de base de datos
// SQLDropTable() string //sql de eliminación de tabla
func NewOperationDB(db *sql.DB, orm OrmAdapter) *operation {
	return &operation{
		DB:         db,
		OrmAdapter: orm,
	}
}
