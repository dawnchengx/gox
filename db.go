package gox

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

var (
    DBtype = "mysql"
    DBconfig = ""
)

func DBconn() *sql.DB {
   db, err := sql.Open(DBtype, DBconfig)
    if err != nil{
        panic(err.Error())
    }
    return db
}
