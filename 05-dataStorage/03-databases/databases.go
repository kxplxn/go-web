package _030503_databases_

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func Demo() {
	fmt.Println("\n030503 Data Storage: Storing and Retrieving Data from a SQL Database")

	db, _ := sql.Open("mysql", "root:P@ssw0rd!@tcp(127.0.0.1:3306)/sys")
	defer db.Close()

	db.Query("CREATE TABLE PERSON (first_name varchar(50), last_name varchar(50))")
	db.Query("INSERT INTO PERSON (first_name, last_name) VALUES ('Bob', 'Barker'), ('Betty', 'White')")

	dataset, _ := db.Query("SELECT * FROM PERSON")
	for dataset.Next() {
		var person Person
		dataset.Scan(&person.FirstName, &person.LastName)
		fmt.Println(person)
	}
}
