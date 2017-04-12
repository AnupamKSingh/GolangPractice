package main
import (
	"fmt"
)

func main () {
	x := []int {1,2,3}
	y := make ([]int,3,5)
	copy (y,x)
	fmt.Println ("Length of y",len(y),"Capacity of y", cap (y),"\n",y,"\n",x)
//	x1 := make ([]string,3,4)
	var x1 []string
	x1 = make ([]string,3,4)
	x1[0]="Anupam"
	x1[2]="Singh"
	x1 =append (x1,"is","good","boy")
	fmt.Println (x1, len (x1), cap (x1))
}
