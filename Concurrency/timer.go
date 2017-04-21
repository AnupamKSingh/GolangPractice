package main
import (
	"fmt"
	"time"
)

func go_func(t time.Duration, i int) <-chan int {
	fmt.Printf("I am in for %d\n", i)
	c := make(chan int)
	go func()	{
		time.Sleep(t)
		c<-i
	}()
	return c
}

func main()	{
	fmt.Println("You are in main")
	for i:=0;i<20;i++ {
		c := go_func(time.Second, i)
		fmt.Printf("Value from channel in main %d",<-c)
	}
}
