package main
import (
	"fmt"
	"sync"
	"time"
)


type order struct {
	dish string  /*Dish Name*/
	num int /*Order Number*/
	duration time.Duration /*Time taken to cook the order*/
}

type chef struct {
	station int /*Station number on which Chef is working*/
	name string /*Name of the Chef*/
}

func newChef(name string, station int) *chef {
	return &chef{
			station:station,
			name:name,}
}

func (c *chef) cook (o *order) {
	fmt.Println("Cooking started", c.name , o.dish)
	time.Sleep(3 * time.Second)
	fmt.Println("Cooking finished", c.name , o.dish)
	c.rest()
}
func (c *chef) rest() {
	fmt.Println("Cook is at rest", c.name)
	time.Sleep(1 * time.Second)
	fmt.Println("Cook is free", c.name)
}

func main() {
	chefs := []*chef{newChef("Anupam", 1), newChef("Deepoo", 1)}
	orders := []*order{
			{"Veg1", 1, 1000 * time.Millisecond},
			{"Veg2", 1, 1000 * time.Millisecond},
			{"Veg3", 1, 1000 * time.Millisecond},
			{"Veg4", 1, 1000 * time.Millisecond},
			{"Veg5", 1, 1000 * time.Millisecond},
			{"Non-Veg1", 2, 2000 * time.Millisecond},
			{"Non-Veg2", 2, 2000 * time.Millisecond},
			{"Non-Veg3", 2, 2000 * time.Millisecond},
			{"Non-Veg4", 2, 2000 * time.Millisecond},
			{"Non-Veg5", 2, 2000 * time.Millisecond},
	}
	startTime := time.Now()
	fmt.Println("Started all orders at ",startTime)
	fireOrders(chefs, orders)
	stopTime := time.Now()
	fmt.Println("Completed All orders at ", stopTime)
}

func fireOrders(chefs []*chef, orders []*order) {
	orderChannel := make (chan *order)
	wg := &sync.WaitGroup{}
	for _, c := range chefs {
		wg.Add(1)
		go func (c *chef) {
			for o := range orderChannel {
				c.cook(o)
			}
			wg.Done()
		}(c)
	}
	for _, order := range orders {
		orderChannel <- order
	}
	close(orderChannel)
	wg.Wait()
}


