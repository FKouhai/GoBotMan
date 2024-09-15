package discover

import (
	"log"
	"net"
	"time"
)
// Discover should test against the config if a station is listening it should add it to the database
// if it is not -> mark it as a dead node
// Discover is the first step that populates the db
func discover (ipAddr string, tcpPort string) {
  tcpaddr := ipAddr + ":" + tcpPort
  log.Println(tcpaddr)
  _, err := net.DialTimeout("tcp",tcpaddr, 5*time.Second)
  if err != nil {
    log.Println(err)
  }
    discovered := ipAddr + " is a listening station"
    log.Println(discovered)
}

// HeartBeat should ping the stations every 500 ms
// if there is a station that doesnt reply mark the node as dead
// if there is a station that replies and was marked as a dead node
// mark it as alive
// HeartBeat should be run as a go routine
func HeartBeat(ipAddr []string, tcpPort string)(){
  log.Println("Checking hearbeat")
  ticker := time.NewTicker(5 * time.Second)
  go func() {
    for _,srv := range ipAddr {
      for range ticker.C{
        discover(srv,tcpPort)
      }
    }
  }()
}
