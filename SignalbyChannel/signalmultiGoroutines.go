/* This code runs multiple goroutines and 
   signals all of them by closiing the channel */
package main
import (
	"fmt"
	"time"
)

func anu (channel chan bool, goroutine int) {
	for {
		select {
			case <-channel:
				fmt.Printf("Channel is closed for goroutine => %d\n", goroutine)
				return
			default:
				fmt.Printf("We are in goroutine => %d\n", goroutine)
				time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	channel := make (chan bool)
	i := 1
	for i < 5 {
		go anu(channel, i)
		i ++
	}
	time.Sleep(5 * time.Second)
	close(channel)
	time.Sleep(2 * time.Second)
}

