package db;

import (
	"fmt"
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
    "time"
    "os"
);

var db *sql.DB

func InitializeConnection()  (*sql.DB,error){
   dBConnection, err := sql.Open("mysql", os.Getenv("db_user_name")+":"+os.Getenv("db_password")+"@(localhost:"+os.Getenv("db_port")+")/"+os.Getenv("db_name"))
   // connection create error
    if err != nil {
        return nil,err
    }
    // check db server state
    err = dBConnection.Ping()
    if err != nil {
        fmt.Println("Ping Failed!!")
        return nil, err
    }
    db = dBConnection
    dBConnection.SetMaxOpenConns(10)
    dBConnection.SetMaxIdleConns(5)
    dBConnection.SetConnMaxLifetime(time.Second * 10)
    // defer db.Close()
    return db, nil;
}

func GetMySQLConnection() (*sql.DB, error) {
    if(db ==nil){
        return InitializeConnection()
    }
    return db, nil;
}