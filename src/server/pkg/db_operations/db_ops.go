package dboperations

import (
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "github.com/rushteam/gosql"
  "os"
)

type StationModel struct {
  ID int `db:"id"`
  Station string `db:"station"`
}

func (u *StationModel) TableName() string {
  return "stations"
}

func createDB(name string, user string, password string, dbHost string) *sql.DB{
  connString := user + ":" + password + "tcp(" + dbHost + ")/"
  db, err := sql.Open("mysql", connString)
  if err != nil {
    panic(err)
  }
  defer db.Close()
  _, err := db.Exec("CREATE DATABASE IF NOT EXISTS "+name)
  if err != nil {
    panic(err)
  }
  return db
  db.Close()
}
