package server

import (
	"log"
	"os"
	"server/pkg/config"
	"server/pkg/db_operations"
)
var db dboperations.DBConn
var station dboperations.StationModel

func main() {
  cfg_file := os.Getenv("CFG_FILE")
  cfg, err := config.NewConfig(cfg_file)
  if err != nil {
    log.Println(err)
  }
  db.User = cfg.User
  db.Password = cfg.Password
  db.DbHost = cfg.DbHost
  station.DBConn = db
  station.Station = cfg.Agents
  station.CreateDB()
}
