package main
import (
	"fmt"
	"time"
)

func anu(channel chan bool) {
	timeout := time.After(5 * time.Second)
	for {
		select {
			case <-timeout:
				fmt.Println("Timeout happened in Goroutine")
				channel <- true
				return
			default:
				fmt.Println("Anupam")
				time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	channel := make(chan bool)
	fmt.Println("Main Begins")
	go anu(channel)
	<-channel
	fmt.Println("Main Exits")
}
