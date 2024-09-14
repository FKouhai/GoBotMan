package dboperations

import (
  "database/sql"
  "log"
  _ "github.com/go-sql-driver/mysql"
)
type DBConn struct {
  User string
  Password string
  DbHost string
}
type StationModel struct {
  DBConn
  Station []string `db:"station"`
  Alive bool `db:"alive"`
}

func (u *StationModel) TableName() string {
  return "stations"
}
func (u *StationModel) CreateDB() *sql.DB{
  db, err := u.DbConn()
  if err != nil {
    log.Fatalln(err)
  }
  defer db.Close()
  _, err = db.Exec("CREATE DATABASE IF NOT EXISTS "+ u.TableName())
  if err != nil {
    log.Fatalln(err)
  }
  return db
}
func (u *StationModel) DbConn() (*sql.DB, error) {
  user := u.User
  password := u.Password
  dbHost := u.DbHost
  connString := user + ":" + password + "tcp(" + dbHost + ")/"
  db, err := sql.Open("mysql", connString)
  if err != nil {
    log.Println(err)
    return nil,err
  }
  return db,nil
}

func(u *StationModel) QueryDB() *sql.Rows{
  db, err := u.DbConn()
  if err != nil {
    log.Println(err)
  }
  stations,err := db.Query("SELECT * FROM STATIONS")
  defer db.Close()
  if err != nil {
    log.Println(err)
  }
  return stations
}
func (u *StationModel) InsertDB() error {
  db, err := u.DbConn()
  if err != nil {
    log.Println(err)
    return err
  }
  insert_q := "INSERT INTO " + u.TableName() 
  values := ` (station, alive) VALUES ($1, $2)`
  db.Begin()
  db.Exec(insert_q + values, u.Station,u.Alive)
  defer db.Close()
  return nil
}
