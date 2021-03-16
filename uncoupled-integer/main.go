package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func findUncoupled(c []int64) int64 {
	ret := int64(0)

	for _, n := range c {
		ret = ret ^ n
	}

	return ret
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	cTemp := strings.Split(strings.TrimSpace(readLine(reader)), ", ")

	var c []int64

	for i := 0; i < len(cTemp); i++ {
		cItem, err := strconv.ParseInt(cTemp[i], 10, 64)
		checkError(err)
		c = append(c, cItem)
	}

	// Print the uncoupled integer given the slice of integers c
	u := findUncoupled(c)

	fmt.Fprintf(writer, "%d", u)

	writer.Flush()
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
