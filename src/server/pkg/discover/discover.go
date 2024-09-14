package discover

import (
	"log"
	"net"
	"time"
)

// Discover should test against the config 
// if a station is listening it should add it to the database
// if it is not -> dont add it
func Discover (ipAddr net.IP, tcpPort string) (string) {
  tcpaddr := ipAddr.String() + ":" + tcpPort
  _, err := net.DialTimeout("tcp",tcpaddr, 5*time.Second)
  if err != nil {
    log.Println(err)
    return (ipAddr.String() + " is not a listening station")
  }else {
    discovered := ipAddr.String() + " is a listening station"
    return discovered
  }

}

// HeartBeat should ping the stations every 500 ms
// if there is a station that doesnt reply back remove it from database
// if there is  a station that replies and isnt on the db adds it
func HeartBeat(ipAddr net.IP, tcpPort string)(){
  log.Println("Checking hearbeat")
  ticker := time.NewTicker(500 * time.Millisecond)
  done := make(chan bool)
  go func() {
    select {
    case <- done:
      return
    case t := <- ticker.C:
      log.Println("Tick at", t)
      tcpAddr := ipAddr.String() + ":" + tcpPort
      net.Dial("tcp", tcpAddr) 
  }
  }()
}
