package main
import (
	"fmt"
_	"sync"
)

//var wg sync.WaitGroup

func main () {
	anu:= make (chan int)
	dee := make (chan int,3)
//	wg.Add(2)	
//	defer wg.Wait()
	go func () {
		fmt.Println ("Anupam")
		fmt.Println (<-anu)
		close(anu)
	}()
	go func () {
		fmt.Println ("Deepoo")
		dee <- 1
		dee <- 2
		dee <- 3
	}()
	anu <- 10
	fmt.Println ("In Main")
	fmt.Println (<-dee)
	n,ok := <-dee
	fmt.Println (n,ok)
	fmt.Println (<-dee)
	close (dee)
}
