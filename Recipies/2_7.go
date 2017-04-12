package main
import (
	"fmt"
)

func main () {
	fmt.Println ("Let's start Panicking")
	panicFunction ()
	fmt.Println ("Panic Ends here")
}

func panicFunction () {
	defer fmt.Println ("Should be Last to display")
	defer func () {
			if e := recover(); e != nil {
				fmt.Println ("Panic is catched")
				fmt.Println ("From Recovery",e)
			}
	}()
	defer func () {
			fmt.Println ("Should be First to display")
	}()
	panic ("Now it's time to panic")
	fmt.Println ("Should never execute")
}
