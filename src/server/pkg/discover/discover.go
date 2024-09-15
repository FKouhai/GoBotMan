package discover

import (
	"log"
	"net"
	"time"
)

// Discover should test against the config 
// if a station is listening it should add it to the database
// if it is not -> mark it as a dead node
// Discover is the first step that populates the db
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
// if there is a station that doesnt reply mark the node as dead
// if there is a station that replies and was marked as a dead node
// mark it as alive
// HeartBeat should be run as a go routine
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
