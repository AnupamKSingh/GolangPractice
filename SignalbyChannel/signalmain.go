/* This code runs one goroutine and signals it for closure,
   by sending data on that channel and then waits on that
   channel so that goroutine may ends its exit process
   and then inally main goroutine ends */
package main
import (
	"fmt"
	"time"
)

func anu(channel chan bool) {
	for {
		select {
			case <-channel:
				fmt.Println("Channel is close is recieved")
				channel <-true
				return
			default:
				fmt.Println("Anupam")
				time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	channel := make (chan bool)
	fmt.Println("Starting the goroutine")
	go anu(channel)
	time.Sleep(5 * time.Second)
	fmt.Println("Channel close is send")
	channel <- true
	fmt.Println("Waiting for response from the goroutine")
	<-channel
	fmt.Println("Goroutine sends the channel closed")
	close(channel)
}
