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
	twos := 0
	threes := 0

	f, err := os.Open("../input")
	check(err)
	defer f.Close()

	in := bufio.NewScanner(f)
	for in.Scan() {
		id := in.Text()
		check(err)

		var count [26]int
		for _, c := range id {
			c -= 'a'
			count[c]++
		}

		for _, n := range count {
			if n == 2 {
				twos++
				break
			}
		}

		for _, n := range count {
			if n == 3 {
				threes++
				break
			}
		}
	}
	check(in.Err())

	fmt.Println(twos * threes)
}
