package main
import (
	"fmt"
	)

func main() {
	var a [10]int
	var n int
	fmt.Println ("Enter number of elements of array");
	fmt.Scanf ("%d",&n);
	fmt.Println ("Enter elements");
	for i:=0;i<n;i=i+1 {
		fmt.Scanf ("%d",&a[i]);
	}
	for i:=0;i<n;i=i+1 {
		fmt.Println (a[i]);
	}
}	
