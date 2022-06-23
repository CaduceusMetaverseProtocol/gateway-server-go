package models

import (
    "database/sql"

    "github.com/gateway-server-go/config"

    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
    "github.com/labstack/gommon/log"
)

var DB *sqlx.DB

func InitDB() {
    DB = NewDB()
}

// DB function
func NewDB() *sqlx.DB {
    db, _ := sqlx.Open("mysql", config.MysqlUser+":"+config.MysqlPassword+"@tcp("+config.MysqlHost+":"+config.MysqlPort+")/"+config.MysqlDatabase+"?charset=utf8")
    err := db.Ping()
    if err != nil {
        log.Panic(err.Error())
    }
    log.Info("Mysql connected.")
    return db
}


// Transaction rollback
func TxRollback(tx *sqlx.Tx) {
    err := tx.Rollback()
    if err != sql.ErrTxDone && err != nil {
        log.Error(err.Error())
    }
}
