package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		panic("give one argument, a number that fits into 4 bytes")
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic("give one argument, a number that fits into 4 bytes")
	}
	fmt.Print(string(rune(n & 0xFF)))
	fmt.Print(string(rune(n & 0xFF00 >> 8)))
	fmt.Print(string(rune(n & 0xFF0000 >> 16)))
	fmt.Print(string(rune(n & 0xFF000000 >> 24)))
}
