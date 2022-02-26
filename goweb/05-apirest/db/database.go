package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const url = "root:mysql4go@tcp(127.0.0.1:3306)/go_db"

var db *sql.DB

func Connect() {
	connection, err := sql.Open("mysql", url)

	if err != nil {
		panic(err)
	}

	db = connection
}

func Close() {
	db.Close()
}

func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

func ExistsTable(table_name string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", table_name)
	rows, err := Query(sql)
	if err != nil {
		fmt.Println(err)
	}
	return rows.Next()
}

func CreateTable(schema string, name string) {
	if !ExistsTable(name) {
		_, err := Exec(schema)
		if err != nil {
			fmt.Println(err)
		}
	}

}

// Limpiar registros
func TruncateTable(table_name string) {
	sql := fmt.Sprintf("TRUNCATE %s", table_name)
	Exec(sql)
}

// Polimorfismo de Exec y Query
func Exec(query string, args ...interface{}) (sql.Result, error) {
	Connect()
	result, err := db.Exec(query, args...)
	Close()
	if err != nil {
		fmt.Println(err)
	}
	return result, err
}

func Query(query string, args ...interface{}) (*sql.Rows, error) {
	Connect()
	rows, err := db.Query(query, args...)
	Close()
	if err != nil {
		fmt.Println(err)
	}
	return rows, err
}
