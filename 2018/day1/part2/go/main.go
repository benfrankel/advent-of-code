package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	freq := 0
	seen := make(map[int]bool)
	deltas := make([]int, 0, 10)

	f, err := os.Open("../input")
	check(err)
	defer f.Close()

	in := bufio.NewScanner(f)

	for in.Scan() {
		if seen[freq] {
			fmt.Println(freq)
			return
		}

		seen[freq] = true

		token := in.Text()
		delta, err := strconv.Atoi(token)
		check(err)

		deltas = append(deltas, delta)
		freq += delta
	}

	check(in.Err())

	for {
		for _, delta := range deltas {
			if seen[freq] {
				fmt.Println(freq)
				return
			}

			seen[freq] = true
			freq += delta
		}
	}
}
