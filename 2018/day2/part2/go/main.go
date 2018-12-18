package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	seen := make(map[struct {
		int
		string
	}]bool)

	f, err := os.Open("../input")
	check(err)
	defer f.Close()

	in := bufio.NewScanner(f)
	for in.Scan() {
		id := in.Text()
		check(err)

		for k := 0; k < len(id); k++ {
			s := id[:k] + id[k+1:]

			sample := struct {
				int
				string
			}{k, s}

			if seen[sample] {
				fmt.Println(s)
				return
			}

			seen[sample] = true
		}
	}
	check(in.Err())
}
