// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	dupFiles := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, dupFiles)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, dupFiles)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			for _, filename := range dupFiles[line] {
				fmt.Println("   " + filename)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int, dupFiles map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		dupFiles[line] = append(dupFiles[line], f.Name())
	}
	// NOTE: ignoring potential errors from input.Err()
}
