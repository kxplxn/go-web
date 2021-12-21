package _030504_sqlMappers_

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
}

func Demo() {
	fmt.Println("\n030504 Data Storage: Using SQL Mappers")

	db, _ := sqlx.Open("mysql", "root:P@ssw0rd!@tcp(127.0.0.1:3306)/sys")
	defer db.Close()

	db.Query("CREATE TABLE PERSON (first_name varchar(50), last_name varchar(50))")
	db.Query("INSERT INTO PERSON (first_name, last_name) VALUES ('Bob', 'Barker'), ('Betty', 'White')")

	tableRec := []Person{}
	db.Select(&tableRec, "SELECT * FROM PERSON")
	fmt.Println(tableRec)
}
