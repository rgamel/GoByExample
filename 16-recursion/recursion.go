package main

import (
	"fmt"
	"example.com/user/GoByExample/16-recursion/tailRecursion"
)

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))
	fmt.Println(tailRecursion.BouncingBall(3, 0.66, 1.5))
}