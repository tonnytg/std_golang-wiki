package crud

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func Create(host string) {

	// conect database and create table
	user := "root"
	password := "teste!123"
	//host := "34.134.48.26"
	port := 3306
	database := "mysql-db"

	fmt.Println("Start connection to database")

	connection := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, database)
	fmt.Println("Conexão:", connection)

	db, err := sql.Open("mysql", connection)
	//db, err := sql.Open("mysql", "root:teste!123@tcp(34.136.118.226:3306)/mysql-db")
	if err != nil {
		panic(err)
	}
	fmt.Println("Conexão:", "ok")

	// create table
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS info (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255), age INT, occupation VARCHAR(255))")
	if err != nil {
		panic(err)
	}
	// parse result
	fmt.Printf("create table ok at %s\n", time.Now().Format("2006-01-02 15:04:05"))

	// insert data
	// insert "Sean", 23, "Content Creator"
	_, err = db.Exec("INSERT INTO info (name, age, occupation) VALUES (?, ?, ?)", "Sean", 23, "Content Creator")
	if err != nil {
		panic(err)
	}
	// insert "Emily", 34, "Cloud Engineer"
	_, err = db.Exec("INSERT INTO info (name, age, occupation) VALUES (?, ?, ?)", "Emily", 34, "Cloud Engineer")
	if err != nil {
		panic(err)
	}
	// inser "Rocky", 40, "Event coordinator"
	_, err = db.Exec("INSERT INTO info (name, age, occupation) VALUES (?, ?, ?)", "Rocky", 40, "Event coordinator")
	if err != nil {
		panic(err)
	}
	// insert "Kate", 28, "Data Analyst"
	_, err = db.Exec("INSERT INTO info (name, age, occupation) VALUES (?, ?, ?)", "Kate", 28, "Data Analyst")
	if err != nil {
		panic(err)
	}

	// insert "Juan", 51, "Program Manager"
	_, err = db.Exec("INSERT INTO info (name, age, occupation) VALUES (?, ?, ?)", "Juan", 51, "Program Manager")
	if err != nil {
		panic(err)
	}

	// insert "Jennifer", 32, "Web Developer"
	_, err = db.Exec("INSERT INTO info (name, age, occupation) VALUES (?, ?, ?)", "Jennifer", 32, "Web Developer")
	if err != nil {
		panic(err)
	}
	fmt.Printf("insert into table ok at %s\n", time.Now().Format("2006-01-02 15:04:05"))
	defer db.Close()
}
