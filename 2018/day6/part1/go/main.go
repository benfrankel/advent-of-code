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
	row int
	col int
}

func Parse(filename string) []point {
	re := regexp.MustCompile(`(\d+), (\d+)`)
	points := make([]point, 0, 10)

	f, err := os.Open(filename)
	check(err)
	defer f.Close()

	in := bufio.NewScanner(f)
	for in.Scan() {
		line := in.Text()
		match := re.FindStringSubmatch(line)

		row, err := strconv.Atoi(match[1])
		check(err)

		col, err := strconv.Atoi(match[2])
		check(err)

		p := point{row, col}
		points = append(points, p)
	}
	check(in.Err())

	return points
}

func BuildGrid(points []point) [][]int {
	// TODO
}

func main() {
	points := Parse("../input")
	grid := BuildGrid(points)
}
