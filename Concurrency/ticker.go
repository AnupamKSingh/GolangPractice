package main
import (
	"fmt"
	"time"
)

func anu(channel chan bool) {
	timeout := time.After(5 * time.Second)
	ticker := time.Tick(1 * time.Second)
	for {
		select {
			case <-ticker:
				fmt.Println("Ticker is called")
			case <-timeout:
				fmt.Println("Timeout is called")
				return
		}
	}
}

func main() {
	channel := make(chan bool)
	fmt.Println("Main Begins")
	go anu(channel)
	time.Sleep(6 * time.Second)
	fmt.Println("Main Ends")
}
