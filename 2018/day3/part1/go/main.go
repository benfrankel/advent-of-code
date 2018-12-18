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

	f, err := os.Open("../input")
	check(err)
	defer f.Close()

	in := bufio.NewScanner(f)
	re := regexp.MustCompile("^#\\d+ @ (\\d+),(\\d+): (\\d+)x(\\d+)$")
	for in.Scan() {
		line := in.Text()
		match := re.FindStringSubmatch(line)
		
		x, err := strconv.Atoi(match[1])
		check(err)

		y, err := strconv.Atoi(match[2])
		check(err)

		w, err := strconv.Atoi(match[3])
		check(err)

		h, err := strconv.Atoi(match[4])
		check(err)

		for i := x; i < x + w; i++ {
			for j := y; j < y + h; j++ {
				count[point {i, j}]++
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
