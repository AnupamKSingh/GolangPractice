package main
import (
	"fmt"
	"time"
)

func ping_pong(ch chan int, i int)	{
	for {
		ball := <-ch
		fmt.Printf("Value of ball is %d in %d\n", ball, i)
		ball ++
		time.Sleep(100 * time.Millisecond)
		ch <- ball
	}
}

func main()	{
	ch := make(chan int)
	for i:= 1;i<100;i++ {
		go ping_pong(ch, i)
	}
	ch <- 0
	time.Sleep(1 * time.Second)
	ball := <-ch
	fmt.Printf("Value of ball is %d\n in main", ball)
}
