package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "1", "2"
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = "1"
		fmt.Println(os.Args[1:])
	}
	fmt.Println(s)
}