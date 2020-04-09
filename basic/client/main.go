package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
	dburl = "root:111111@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local&allowAllFiles=true"
	sqlString = `LOAD DATA INFILE "./electricData.csv" INTO TABLE electrion FIELDS TERMINATED BY ',' ENCLOSED BY '"' LINES TERMINATED BY '\n'  IGNORE 1 LINES`
)

func init() {
	var err error
	db, err = sql.Open("mysql", dburl)
	if err != nil {
		panic(err)
	}
}

func initTable() {
	
}

func main() {
	_, err := db.Exec(sqlString)
	if err != nil {
		fmt.Println(err)
	}
}