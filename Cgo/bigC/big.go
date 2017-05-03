package main
/*
#include<big.h>
#include<stdio.h>
void big() {
	printf("Anupam");
}
void anu() {
	big();
}
#cgo LDFLAGS: -L.  -lbig
*/
import "C"

func main() {
	C.big()
//	C.anu()
}
