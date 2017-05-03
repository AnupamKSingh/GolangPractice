/* This code run one goroutine and signals 
   the channel for closure by closing the channel */
package main
import (
	"fmt"
	"time"
)

func anu(channel chan bool) {
	for {
		select {
			default:
					fmt.Println("Anupam")
					time.Sleep(1 * time.Second)
			case <- channel:
				fmt.Println("Channel closed")
				return
		}
		}
}

func main() {
	channel := make (chan bool)
	fmt.Println("We are in Main")
	go anu (channel)
	time.Sleep(5 * time.Second)
	close(channel)
	time.Sleep(2 * time.Second)
}
