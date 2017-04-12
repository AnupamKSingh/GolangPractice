package main
import (
	"fmt"
	"syscall"

	"github.com/vishvananda/netlink"
)

func main()	{
	var link netlink.Link
	fmt.Println ("Lets have a try")
	list, err := netlink.AddrList (&link, syscall.AF_INET)
	if err != nil	{
		fmt.Println ("Error Happened ")
		fmt.Println (err)
	}	else	{
		fmt.Println ("Your success")
		fmt.Println (list)
	}
}
		
