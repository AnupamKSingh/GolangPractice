package main 
import (
	"fmt"
)
/* Once can Omit the use of global variable (Package level access)
   Will not through compiler error */
var a int

func main () {
	var b int
	b = 12
	/* But if local variables is not used but declared
	   Compiler will throw error */
//	fmt.Println(a)
	fmt.Println(b)
}
