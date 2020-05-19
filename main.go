package geeorm

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	db, _ := sql.Open("sqlite3", "gee.db")

	defer func() { _ = db.Close() }()

	_, _ = db.Exec("DROP TABLE IF EXISTS USER;")
	_, _ = db.Exec("CREATE TABLE User(Name TEXT)")
	result, err := db.Exec("INSERT INTO User(`Name`) VALUES (?), (?)", "Tom", "Jack")
	if err == nil {
		affected, _ := result.RowsAffected()
		log.Printf("affected rows: %d\n", affected)
	}

	row := db.QueryRow("SELECT Name FROM User LIMIT 1")
	var name string
	if err := row.Scan(&name); err == nil {
		log.Println(name)
	}
}
