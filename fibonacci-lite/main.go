package main

import "fmt"

func main() {
	n := readInt()
	fmt.Println(falcFibonnaciLite(n))
}

func falcFibonnaciLite(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	return falcFibonnaciLite(n-2) + falcFibonnaciLite(n-1)
}

func readInt() int {
	var n int
	_, err := fmt.Scanf("%d", &n)

	if err != nil {
		panic(err)
	}

	return n
}
