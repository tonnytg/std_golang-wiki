package gosql

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type Names struct {
	id string `db:"id"`
	names string `db:"names"`
}

func GoSQL() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	table := `CREATE TABLE names (
		"id" string,
		"name" string
	);`

	stmt, err := db.Prepare(table)
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec()
	if err != nil {
		panic(err)
	}

	insert := `insert into names values("1","bar")`
	stmt, err = db.Prepare(insert)
	if err != nil {
		panic(err)
	}
	stmt.Exec()

	insert = `insert into names values("2","foo")`
	stmt, err = db.Prepare(insert)
	if err != nil {
		panic(err)
	}
	stmt.Exec()

	var name Names
	rows, _ := db.Query("SELECT * FROM names")
	for rows.Next() {
		rows.Scan(&name.id, &name.names)
		fmt.Printf("%v\n", name)
	}
}
