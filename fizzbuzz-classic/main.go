package main

import "fmt"

func main() {
	number := readInt()

	for i := 0; i < number; i++ {
		div3 := ((i + 1) % 3) == 0
		div5 := ((i + 1) % 5) == 0
		if div3 && div5 {
			fmt.Println("FizzBuzz")
		} else if div3 {
			fmt.Println("Fizz")
		} else if div5 {
			fmt.Println("Buzz")
		} else {
			fmt.Printf("%d\r\n", i+1)
		}
	}
}

func readInt() int {
	var n int
	_, err := fmt.Scanf("%d", &n)

	if err != nil {
		panic(err)
	}

	return n
}
