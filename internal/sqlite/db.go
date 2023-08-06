package sqlite

import (
	"database/sql"
	"log"
	// Import the SQLite driver
	_ "modernc.org/sqlite"
)

func Connect(dbFile string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", dbFile)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Exec(db *sql.DB, query string, args ...any) (sql.Result, error) {
	result, err := db.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func ExecPrepare(stmt *sql.Stmt, args ...any) (sql.Result, error) {
	result, err := stmt.Exec(args...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func Query(db *sql.DB, query string, args ...any) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		log.Printf("error with db.Query!")
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(rows)
	return rows, nil
}
