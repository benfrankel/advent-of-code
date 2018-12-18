package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type point struct {
	x int
	y int
}

func main() {
	count := make(map[point]int)
	re := regexp.MustCompile("^#\\d+ @ (\\d+),(\\d+): (\\d+)x(\\d+)$")

	f, err := os.Open("../input")
	check(err)
	defer f.Close()

	in := bufio.NewScanner(f)
	for in.Scan() {
		line := in.Text()
		match := re.FindStringSubmatch(line)

		l, err := strconv.Atoi(match[1])
		check(err)

		t, err := strconv.Atoi(match[2])
		check(err)

		w, err := strconv.Atoi(match[3])
		check(err)
		r := l + w

		h, err := strconv.Atoi(match[4])
		check(err)
		b := t + h

		for i := l; i < r; i++ {
			for j := t; j < b; j++ {
				count[point{i, j}]++
			}
		}
	}
	check(in.Err())

	num := 0
	for _, n := range count {
		if n >= 2 {
			num++
		}
	}

	fmt.Println(num)
}
