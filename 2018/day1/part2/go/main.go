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

func Append(slice []int, n int) []int {
	l := len(slice)

	if l >= cap(slice) {
		newSlice := make([]int, cap(slice)*2)
		copy(newSlice, slice)
		slice = newSlice
	}

	slice = slice[0 : l+1]
	slice[l] = n

	return slice
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

		deltas = Append(deltas, delta)
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
