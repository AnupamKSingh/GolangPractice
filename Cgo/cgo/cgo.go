package main
/*
#include<stdio.h>
#cgo LDFLAGS: -L. -lcgo

void PrintAnu();
void PrintAnu() {
	printf("\nAnupam\n");
}
*/
import "C"

func main() {
	C.PrintAnu()
}
