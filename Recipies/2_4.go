package main
import (
	"fmt"
)

func main () {
	var x map [string]string
	x =make (map[string]string)
	x["Anupam"] = "Good"
	x["Deepoo"] = "Bad"
	for key,value := range x {
		fmt.Println (key,value)
	}
	y := map[int]string {
		1:"Anupam",
		3:"Deepoo",
	}
	for key,value := range y {
		fmt.Println (key,value)
	}
	/* We should use range for printing maps , since Range will show all nonzero values */
	for i :=0;i <len(y);i++ {
		fmt.Println(i, y[i])
	}
	z := map[string]int {
		"Anupam":1,
		"Deepoo":23,
	}
	for key,value := range z {
		fmt.Println (key,value)
	}
	if value,ok := z["Anupam"];ok {
		fmt.Println ("Found",value)
	}else{
		fmt.Println ("Not Found")
	}
}
