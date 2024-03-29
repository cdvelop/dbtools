package dbtools

import "database/sql"

type dboAdapter interface {
	dbAdapter
	OrmAdapter
}

type dbAdapter interface {
	Open() *sql.DB
}

type OrmAdapter interface {
	//ej postgres:"$1", sqlite: "?"
	PlaceHolders(index ...uint8) string
	DeleteDataBase()
	// SQLTableInfo() string //sql como obtiene la base de datos el nombre de la tabla
	SQLTableInfo() string
	// SQLColumName() string //sql como se llama a la columna en el motor de base de datos
	SQLColumName() string
	// SQLDropTable() string //sql de eliminación de tabla
	SQLDropTable() string
	//ej: pg: DROP TABLE IF EXISTS $1 CASCADE; sqlite:
	DropTable() string
	//ej: pg: SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_schema = 'public' AND table_name = $1)
	SQLTableExist() string
}
