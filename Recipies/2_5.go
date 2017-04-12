package main
import (
	"fmt"
	"io/ioutil"
	"os"
)

func main () {
	f,err := os.Open ("a.txt")
	if err != nil {
		fmt.Println ("Error Happened")
	}
	defer f.Close()
	str,err :=ioutil.ReadAll (f)
	fmt.Println (err)
	fmt.Println (string(str))
}
	
