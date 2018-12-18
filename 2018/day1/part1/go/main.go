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
	f, err := os.Open("../input")
	check(err)
	defer f.Close()

	in := bufio.NewScanner(f)
	freq := 0
	for in.Scan() {
		token := in.Text()
		delta, err := strconv.Atoi(token)
		check(err)

		freq += delta
	}

	fmt.Println(freq)
	check(in.Err())
}
