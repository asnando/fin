package common

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
}

func (d Database) createConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "db/fin.db")

	if err != nil {
		panic(err)
	}

	return db
}

func (d Database) Query(tableName string) *sql.Rows {
	db := d.createConnection()

	defer db.Close()

	q := fmt.Sprintf("SELECT * FROM %s", tableName)

	rows, err := db.Query(q)

	if err != nil {
		panic(err)
	}

	return rows
}

func (d Database) RawQuery(q string) *sql.Rows {
	db := d.createConnection()

	defer db.Close()

	rows, err := db.Query(q)

	if err != nil {
		panic(err)
	}

	return rows
}

func (d Database) Insert(table string, fields string, data ...interface{}) {
	db := d.createConnection()

	defer db.Close()

	quotes := make([]string, 0)

	for i := 0; i < len(data); i++ {
		quotes = append(quotes, "?")
	}

	quotedValue := strings.Join(quotes, ",")

	q := fmt.Sprintf("INSERT INTO %s(%s) VALUES(%s);", table, fields, quotedValue)
	fmt.Println(q)

	stmt, err := db.Prepare(q)
	Error{}.CheckErr(err)

	_, err = stmt.Exec(data...)
	Error{}.CheckErr(err)

	fmt.Println("Successfully inserted into", table)
}
