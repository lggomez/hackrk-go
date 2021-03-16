package main

import (
	"bufio"
	"fmt"
	"io"
	"math/big"
	"os"
	"strconv"
	"strings"
)

var cache = map[uint64]*big.Int{}

var bigIntZero = big.NewInt(int64(0))
var bigIntOne = big.NewInt(int64(1))

func fibonacciReturns(n uint64) *big.Int {
	if val, found := cache[n]; found {
		return val
	}

	if n == 0 {
		cache[0] = bigIntZero
		return cache[0]
	} else if n == 1 {
		cache[1] = bigIntOne
		return cache[1]
	}

	fib := new(big.Int).Set(fibonacciReturns(n - 1))
	cache[n] = fib.Add(fib, fibonacciReturns(n-2))

	return cache[n]
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	var c []int64

	for {
		v := readLine(reader)
		if v == "" {
			break
		}
		firstMultipleInput := strings.Split(strings.TrimSpace(v), " ")
		mTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)

		checkError(err)
		c = append(c, mTemp)
	}

	f := []*big.Int{}
	for _, number := range c {
		f = append(f, fibonacciReturns(uint64(number)))
	}

	fmt.Fprintf(writer, "%s", printSlice(f))

	writer.Flush()
}

func printSlice(numbers []*big.Int) string {
	numberStr := ""
	for i, n := range numbers {
		if i == 0 {
			numberStr = fmt.Sprintf("%d", n)
		} else {
			numberStr = fmt.Sprintf("%s\n%d", numberStr, n)
		}
	}
	return numberStr
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
