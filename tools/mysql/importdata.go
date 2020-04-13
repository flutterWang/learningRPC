package mysql

import (
	"database/sql"
	"errors"
	"fmt"
)

var (
	sqlString = `LOAD DATA INFILE "%v" INTO TABLE %v FIELDS TERMINATED BY ',' ENCLOSED BY '"' LINES TERMINATED BY '\n'  IGNORE 1 LINES`
)

// ImportData -
func ImportData(db *sql.DB, filename string, name string) error {
	execString := fmt.Sprintf(sqlString, filename, "electrion")
	_, err := db.Exec(execString)

	if err != nil {
		return errors.New("import data fail")
	}

	return nil
}
