package main
import (
	"fmt"
	"sync"
	"math/rand"
	"time"
	"runtime"
)

var wg sync.WaitGroup
func main () {
	fmt.Println (runtime.GOMAXPROCS(8))
	wg.Add(2)
	go add()
	go multiply()
	wg.Wait()
}

func add () {
	defer wg.Done()
	for i:= 0;i<10;i++ {
		sleep := rand.Int63n (1000)
		time.Sleep (time.Duration(sleep) * time.Millisecond)
		for j:=0;j<10;j++ {
			fmt.Println("Add ", i," ",j," ",i+j)
		}
	}
}

func multiply () {
	defer wg.Done()
	for i:= 0;i<10;i++ {
		sleep := rand.Int63n (1000)
		time.Sleep (time.Duration(sleep) * time.Millisecond)
		for j:=0;j<10;j++ {
			fmt.Println("Multiply ", i," ",j," ",i*j)
		}
	}
}
