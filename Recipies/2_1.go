package main

import (
	"fmt"
)

func main () {
	var b,c int
	fmt.Println ("Enter two Numbers")
	fmt.Scanf ("%d",&b)
	fmt.Scanf ("%d",&c)
	a := add (b,c)
	fmt.Println ("Output is ",a)
	var sa,sb = "Anupam","Singh"
	sc := combine (sa,sb)
	fmt.Println ("Output is ",sc)
	sum := sum ("Anupam",1,2,3,4)
	fmt.Println ("Sum Output is ",sum)
	x,y := swap (2,3)
	fmt.Printf ("x = %d, y = %d",x,y)
}

func add (a,b int) (result int) {
	result =a +b
	return
}

func combine (a,b string) (r string) {
	r = a+b
	return
}

func sum (str string,a ...int) (s int) {
	s =0
	fmt.Println ("Welcome ",str)
	for _,a := range a {
		s = s+a
	}
	return
}

func swap (x,y int ) (a,b int) {
	a,b=y,x
	return
}
