package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(concat(os.Args[1:]), join(os.Args[1:]))
}

func concat(str []string) string {
	s, sep := "", ""
	for _, arg := range str {
		s += sep + arg
		sep = " "
	}
	return s
}

func join(str []string) string {
	return strings.Join(str, " ")
}
