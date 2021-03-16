package main

import (
	"fmt"
)

func parseString(s string) bool {
	if len(s)%2 == 1 {
		return false
	}

	delimiterIndex := map[string]int{
		"(": 0,
		")": 0,
		"[": 0,
		"]": 0,
		"{": 0,
		"}": 0,
	}

	commaBoundaryCheck := func() int {
		return delimiterIndex["("] - delimiterIndex[")"]
	}

	squareBracketBoundaryCheck := func() int {
		return delimiterIndex["["] - delimiterIndex["]"]
	}

	bracketBoundaryCheck := func() int {
		return delimiterIndex["["] - delimiterIndex["]"]
	}

	for _, c := range s {
		//fmt.Printf("'%v' - '%v'\n", string(c), string(s[len(s)-i-1]))
		c1 := string(c)
		delimiterIndex[c1] = delimiterIndex[c1] + 1
	}

	return (commaBoundaryCheck() == 0) || (squareBracketBoundaryCheck() == 0) || (bracketBoundaryCheck() == 0)
}

func main() {
	switch parseString(readString()) {
	case true:
		fmt.Println("True")
	case false:
		fmt.Println("False")
	}
}

func readString() string {
	var s string
	_, err := fmt.Scanf("%s", &s)

	if err != nil {
		panic(err)
	}

	return s
}
