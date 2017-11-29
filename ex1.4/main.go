// go run main.go file1.txt file2.txt file3.txt

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, c := range counts {
		all := 0
		for _, n := range c {
			all += n
		}
		if all > 1 {
			fmt.Printf("%d\t%s\n", all, line)
			for f := range c {
				fmt.Printf("\t- %s\n", f)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if _, ok := counts[input.Text()]; !ok {
			counts[input.Text()] = make(map[string]int)
		}
		counts[input.Text()][f.Name()]++
	}
}
