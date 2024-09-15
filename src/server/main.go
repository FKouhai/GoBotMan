package main

import (
	"log"
	"os"
	"server/pkg/config"
	"server/pkg/db_operations"
  "server/pkg/discover"
  "github.com/gin-gonic/gin"

)
var db dboperations.DBConn
var station dboperations.StationModel

// server should have a web API as well that allows to control the stations and returns its status
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
  discover.HeartBeat(station.Station,cfg.Port)
  router := gin.Default()
  log.Println(router.Run())
}
