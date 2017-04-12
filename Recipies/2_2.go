package main
import (
	"fmt"
)

func main () {
	s := [4]string {"Anupam","Kumar","Singh"}
	for _,str := range s {
		fmt.Printf ("\n %s",str)
	}
	s = [4]string {1:"Deepoo",3:"Anupam"}
	for _,str := range s {
		fmt.Println (str)
	}
	s1 := [...]string {"Are","you","sure"}
	for _,str := range s1 {
		fmt.Println(str)
	}
}
