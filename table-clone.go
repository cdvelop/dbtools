package dbtools

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/cdvelop/model"
)

// ClonDATABLE copia la data de una tabla a otra nueva
func ClonDATABLE(o dboAdapter, table *model.Object) (err string) {
	const this = "ClonDATABLE"
	db := o.Open()
	defer db.Close()
	// fmt.Printf("Clon Object: %v\n", table.Name)

	// `tx` es una instancia de` * sql.Tx` a través de la cual podemos ejecutar nuestras consultas
	ctx := context.Background()
	tx, e := db.BeginTx(ctx, nil) // Crea un nuevo contexto y comienza una transacción
	if e != nil {
		return this + e.Error()
	}
	defer tx.Rollback()

	err = ClonOneTableInTransaction(o, table, tx, ctx)
	if err != "" {
		tx.Rollback()
		return this + err
	}

	// Finalmente, si no se reciben errores de las consultas, confirme la transacción
	// esto aplica los cambios anteriores a nuestra base de datos
	e = tx.Commit()
	if e != nil {
		tx.Rollback()
		return this + e.Error()
	}

	return ""
}

// ClonOneTableInTransaction copia la data de una tabla a otra nueva
func ClonOneTableInTransaction(o dboAdapter, table *model.Object, tx *sql.Tx, ctx context.Context) (err string) {
	const this = "ClonOneTableInTransaction"
	// fmt.Printf("Clon Object: %v\n", table.Name)
	tableTempName := `tabtemp`

	// 1 renombrar tabla
	_, e := tx.ExecContext(ctx, `ALTER TABLE `+table.ObjectName+` RENAME TO `+tableTempName+`;`)
	if e != nil {
		tx.Rollback()
		return "!!! " + this + e.Error() + " al renombrar tabla " + table.ObjectName
	}
	// fmt.Printf(">>>[1] RENOMBRAR TABLA TABLA: %v\n", tableTempName)

	//2 crear tabla nueva
	sqlNewTable := makeSQLCreaTABLE(table)
	// fmt.Printf(">>>[2] SQL NUEVA TABLA: %v\n", sqlNewTable)
	_, e = tx.ExecContext(ctx, sqlNewTable)
	if e != nil {
		tx.Rollback()
		return this + e.Error()
	}
	//3 seleccionar data anterior
	var oldfield []string
	sqlOldField := fmt.Sprintf(o.SQLTableInfo(), table.ObjectName)
	// fmt.Printf(">>>[3] SELECCIONAR DATA ANTERIOR SQL OLDFIELD: %v\n", sqlOldField)

	// knames, ok = tx.getallOBJ(&q, &ctx)
	knames, err := SelectAll(sqlOldField, o, ctx) //entrega nombre columnas de la tabla
	if err != "" {
		tx.Rollback()
		return this + err
	}

	// fmt.Printf(">>>[4] %v COLUMNAS PARA COPIAR: %v\n", len(knames), knames)

	if len(knames) == 0 {
		tx.Rollback()
		return "!!! " + this + " sin columnas para copiar"
	}

	for _, d := range knames {
		oldfield = append(oldfield, d[o.SQLColumName()])
	}

	var toClone []string //columnas a copiar
	for _, field := range table.Fields {
		for _, oldfield := range oldfield {
			if field.Name == oldfield {
				toClone = append(toClone, oldfield) //agrego las columnas que serán copiadas
				break
			}
		}
	}

	// fmt.Printf(">>>[5] A CLONAR: %v\n", toClone)
	// fmt.Printf(">>> toClone %v\n", toClone)
	//4 copiar data
	c := strings.Join(toClone, ",") //creando un string separado por ,
	// INSERT INTO ciudad (idciudad,nombre) SELECT idciudad,nombre FROM temp
	sqlInsert := fmt.Sprintf("INSERT INTO %v (%v) SELECT %v FROM %v;", table.ObjectName, c, c, tableTempName)
	// fmt.Printf(">>> sql insert %v\n", sqlInsert)
	// fmt.Printf(">>> copiando data %v\n", table.Name)
	_, e = tx.ExecContext(ctx, sqlInsert)
	if e != nil {
		tx.Rollback()
		return "!!! " + this + e.Error() + " al copiar data de " + tableTempName + " a tabla " + table.ObjectName
	}
	// fmt.Printf(">>>[6] DATA COPIADA: %v\n", table.Name)
	// fmt.Printf(">>> data copiada: %v\n", table.Name)

	//5 borrar tabla temporal
	sqlDelete := fmt.Sprintf(o.SQLDropTable(), tableTempName)
	// log.Printf(">> sql droptab : %v", q)
	_, e = tx.ExecContext(ctx, sqlDelete)
	if e != nil {
		tx.Rollback()
		return "!!! " + this + e.Error() + " al borrar tabla temporal " + table.ObjectName
	}

	fmt.Printf(">>> TABLA: %v CLONADA OK\n", table.ObjectName)

	return ""
}
