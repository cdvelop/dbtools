package dbtools

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/cdvelop/model"
)

// ClonDATABLE copia la data de una tabla a otra nueva
func (o operation) ClonDATABLE(table model.Object) bool {
	// fmt.Printf("Clon Object: %v\n", table.Name)

	// `tx` es una instancia de` * sql.Tx` a través de la cual podemos ejecutar nuestras consultas
	ctx := context.Background()
	tx, err := o.DB.BeginTx(ctx, nil) // Crea un nuevo contexto y comienza una transacción
	if err != nil {
		log.Println(err)
		return false
	}

	defer tx.Rollback()

	if !o.ClonOneTableInTransaction(table, tx, ctx) {
		tx.Rollback()
		return false
	}

	// Finalmente, si no se reciben errores de las consultas, confirme la transacción
	// esto aplica los cambios anteriores a nuestra base de datos
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
		return false
	}

	return true
}

// ClonOneTableInTransaction copia la data de una tabla a otra nueva
func (o operation) ClonOneTableInTransaction(table model.Object, tx *sql.Tx, ctx context.Context) bool {
	// fmt.Printf("Clon Object: %v\n", table.Name)
	var ok bool
	tableTempName := `tabtemp`

	// 1 renombrar tabla
	_, err := tx.ExecContext(ctx, `ALTER TABLE `+table.Name+` RENAME TO `+tableTempName+`;`)
	if err != nil {
		log.Printf("!!! error %v al renombrar tabla %v", err, table.Name)
		tx.Rollback()
		return false
	}
	// fmt.Printf(">>>[1] RENOMBRAR TABLA TABLA: %v\n", tableTempName)

	//2 crear tabla nueva
	sqlNewTable := makeSQLCreaTABLE(table)
	// fmt.Printf(">>>[2] SQL NUEVA TABLA: %v\n", sqlNewTable)
	_, err = tx.ExecContext(ctx, sqlNewTable)
	if err != nil {
		tx.Rollback()
		return false
	}
	//3 seleccionar data anterior
	var oldfield []string
	sqlOldField := fmt.Sprintf(o.SQLTableInfo(), table.Name)
	// fmt.Printf(">>>[3] SELECCIONAR DATA ANTERIOR SQL OLDFIELD: %v\n", sqlOldField)

	// knames, ok = tx.getallOBJ(&q, &ctx)
	var knames = make([]map[string]string, 0)
	if knames, ok = SelectAll(sqlOldField, o.DB, ctx); !ok { //entrega nombre columnas de la tabla
		tx.Rollback()
		return false
	}
	// fmt.Printf(">>>[4] %v COLUMNAS PARA COPIAR: %v\n", len(knames), knames)

	if len(knames) == 0 {
		tx.Rollback()
		log.Println("!!! error sin columnas para copiar")
		return false
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
	sqlInsert := fmt.Sprintf("INSERT INTO %v (%v) SELECT %v FROM %v;", table.Name, c, c, tableTempName)
	// fmt.Printf(">>> sql insert %v\n", sqlInsert)
	// fmt.Printf(">>> copiando data %v\n", table.Name)
	_, err = tx.ExecContext(ctx, sqlInsert)
	if err != nil {
		log.Printf("!!! error %v al copiar data de %v a tabla %v", err, tableTempName, table.Name)
		tx.Rollback()
		return false
	}
	// fmt.Printf(">>>[6] DATA COPIADA: %v\n", table.Name)
	// fmt.Printf(">>> data copiada: %v\n", table.Name)

	//5 borrar tabla temporal
	sqlDelete := fmt.Sprintf(o.SQLDropTable(), tableTempName)
	// log.Printf(">> sql droptab : %v", q)
	_, err = tx.ExecContext(ctx, sqlDelete)
	if err != nil {
		log.Printf("!!! error %v al borrar tabla temporal %v", err, table.Name)
		tx.Rollback()
		return false
	}

	fmt.Printf(">>> TABLA: %v CLONADA OK\n", table.Name)
	return true
}
