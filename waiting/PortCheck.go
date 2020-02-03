package waiting

import (
	"fmt"
	"net"
	"time"
)

func NewPortCheck(hostWithPort string) CheckTask {
	return func(timeout time.Duration) bool {
		fmt.Print("Checking port: ", hostWithPort)
		con, err := net.DialTimeout("tcp4", hostWithPort, timeout)
		if con != nil {
			con.Close()
		}
		fmt.Println("", err == nil)
		return err == nil
	}
}
