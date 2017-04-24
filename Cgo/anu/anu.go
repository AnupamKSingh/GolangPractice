package main
/*
#include<stdio.h>
#cgo LDFLAGS: -L. -lanu
void anu();
void anu() {
	int i;
	for (i=1; i<10; i++)
		printf("\nAnupamK\n");
	printf("\nDeepoo\n");
}
*/
import "C"
func main() {
	C.anu()
}
