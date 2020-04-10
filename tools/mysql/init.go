package mysql

import "database/sql"

// InitDB -
func InitDB(url string) *sql.DB {
	db, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}

	return db
}
