package main
import (
	"fmt"
_	"io"
_	"io/ioutil"
_	"bufio"
)

var message string

const CLR_0 = "\x1b[33;1m"
const CLR_N = "\x1b[0m"

func main ()	{
	fmt.Println ("Lets type something")
/*	fmt.Scanf ("%q", &message)
	fmt.Println ("Type is completed")
	fmt.Println (message)
	fmt.Println (message)
	fmt.Println ("Message print Ends")
*/
	mess := "Anupam Kumar Singh"
	fmt.Printf ("%s%s%s\n",CLR_0, mess, CLR_N)
	fmt.Println("\x1b[31;1mHello, World!\x1b[0m")
	fmt.Println("\x1b[33;1mHello, World!\x1b[0m")
	fmt.Println("\x1b[34;1mHello, World!\x1b[0m")
	fmt.Println("\x1b[35;1mHello, World!\x1b[0m")
}
