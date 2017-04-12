// +build linux,386 darwin/amd64
package main
import (
	"syscall"
	"fmt"
)

type NetlinkSocket struct {
    fd  int  
// +build linux,386 darwin,!cgo
    lsa syscall.SockaddrNetlink
}

func main()	{
	s, err := syscall.Socket(0x10, syscall.SOCK_RAW, 0x5) // NFLOG = iptables ULOG
	if err != nil {
		fmt.Println("******  NetLink Error 1:", err)
	}
	// +build linux,386 darwin,!cgo
	ns := &NetlinkSocket{
        fd: s,
    }
	if err := syscall.Bind(s, ns.lsa); err != nil { // If Pid==0, Kernel chooses
		fmt.Println("******  NetLink Error 2:", err)
	}
	var buffer []byte
	n, err := syscall.Read(s, buffer)
	fmt.Println (err)
}
