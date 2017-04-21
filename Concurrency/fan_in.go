package main
import (
	"fmt"
	"time"
)

func producer(ch chan<- int, t time.Duration) {
	i := 1
	for {
		ch <- i
		i ++
		time.Sleep(t)
	}
}

func consumer(ch <-chan int)	{
	for x := range ch {
		fmt.Println(x)
	}
}

func main() {
	chp := make(chan int)
	chc := make(chan int)
	go producer(chp, 100 * time.Millisecond)
	go producer(chp, 250 * time.Millisecond)
	go consumer(chc)
	for i := range chp {
		chc <- i
	}
}
