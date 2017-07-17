package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/Zhanat87/golang-gorilla/util"
	"os"
	"fmt"
)

var Connection *sql.DB

func init() {
	dsn := fmt.Sprintf("%s:%s@/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DATABASE"))
	var err error
	Connection, err = sql.Open("mysql", dsn)
	util.FailOnError(err, "error mysql connection")
	err = Connection.Ping()
	util.FailOnError(err, "error mysql ping")
}
