package main
import (
	"fmt"
)

func main()	{
	fmt.Println("You are in main")
	ch := make(chan int)
	go func()	{
		fmt.Printf("You are in go func")
		ch <- 42
	}()
	fmt.Println("End of go func")
	<-ch
	fmt.Println("End of main")
}

