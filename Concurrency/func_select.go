/* In the below code for select statement though We are hitting the case 1 (anu())
   But the function dee() is also being called and the value is passed through the channel */
package main
import (
	"fmt"
_	"time"
)

func anu() int{
	fmt.Println("Anupam")
	return 0
}

func dee() int {
	fmt.Println("Deepoo")
	return 1
}

func channel(chan1 , chan2 chan int) {
	for {
		select {
			case chan1 <- anu():
				fmt.Println("anu() is called")
				break
			case chan2<- dee():
				fmt.Println("dee() is called")
				break
		}
	}
}

func main() {
	chan1, chan2 := make (chan int), make (chan int)
	fmt.Println("Main Begins")
	go channel(chan1, chan2)
	fmt.Println("Channel started")
	fmt.Printf("Channel 1 recieves = %d\n",<-chan1)
	fmt.Println("Chan1 is recieved")
	fmt.Printf("Channel 2 recieves = %d\n",<-chan2)
	fmt.Println("Chan2 is recieved")
	fmt.Println("Main Exits")
}
