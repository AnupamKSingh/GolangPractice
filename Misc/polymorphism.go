package main
import (
	"fmt"
)

func anu(f func(int) int, xs []int) []int {
		ys := make([]int, len(xs))
		for i, x := range xs {
				ys[i] = f(x)
		}
		return ys
}
func main() {
	fmt.Println(anu(func(x int) int { return x * x }, []int{1, 2, 3}))
}
