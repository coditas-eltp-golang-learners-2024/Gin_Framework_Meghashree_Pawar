package config

import (
    "database/sql"
    "fmt"

    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() error {
    var err error
    dataSourceName := "root:password@tcp(localhost:3306)/api"
    DB, err = sql.Open("mysql", dataSourceName)
    if err != nil {
        return err
    }
    err = DB.Ping()
    if err != nil {
        return err
    }
    fmt.Println("Connected to MySQL database")
    return nil
}

