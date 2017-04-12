package main
import (
	"fmt"
)

const (
	a = iota
	b = iota +10
	c = iota
)

func main () {
	fmt.Println (a)	
	fmt.Println (b)	
	fmt.Println (c)	
}
