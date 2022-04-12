package discover

import (
  "net"
  "fmt"
  "time"
)

func discover (ipAddr IP, tcpPort string) (string) {
  tcpaddr := ipAddr.String() + ":" + tcpPort
  _, err := net.DialTimeout("tcp",tcpaddr, 5*time.Second)
  if err != nil {
    fmt.Println(err)
    return nil
  }else {
    discovered := ipAddr.String() + " is a listening station"
    return discovered
  }

}
