package sql

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

func NewMysql() (db *sql.DB){

	db,err := sql.Open("mysql","root:123456/dami"));
	if err!=nil{
		panic(err)
	}

	return db
}