package main

import (
	"fmt"
	"math/big"
)

func main() {
	n := readInt()
	bigInt := big.NewInt(int64(n))
	fmt.Println(factorial(bigInt))
}

var bigIntZero = big.NewInt(int64(0))
var bigIntOne = big.NewInt(int64(1))

func factorial(n *big.Int) *big.Int {
	if n.Cmp(bigIntZero) == -1 {
		return bigIntZero
	}
	if n.Cmp(bigIntZero) == 0 {
		return bigIntOne
	} else if n.Cmp(bigIntOne) == 0 {
		return bigIntOne
	}

	ret := new(big.Int).Set(n)
	return ret.Mul(ret, factorial(n.Sub(n, bigIntOne)))
}

func readInt() int {
	var n int
	_, err := fmt.Scanf("%d", &n)

	if err != nil {
		panic(err)
	}

	return n
}
