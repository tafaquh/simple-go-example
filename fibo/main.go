package main

import (
	"fmt"
	"os"
	"strconv"
)

func FibonacciRecursion(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciRecursion(n-1) + FibonacciRecursion(n-2)
}

func FibonacciLoop(n int) int {
	f := make([]int, n+1, n+2)
	if n < 2 {
		f = f[0:2]
	}
	f[0] = 0
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}

func main() {
	// Read input
	n := os.Getenv("NUMBER")

	number, _ := strconv.Atoi(n)

	// Run Fibo
	result := FibonacciLoop(number)

	fmt.Printf("\nFibonacchi of %d is %d\n", number, result)

}
