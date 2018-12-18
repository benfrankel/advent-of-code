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

func main() {
	re := regexp.MustCompile("^[(\\d+)-(\\d+)-(\\d+) (\\d+):(\\d+)] (Guard #(\\d+) begins shift|falls asleep|wakes up)$")

	f, err := os.Open("../input")
	check(err)
	defer f.Close()

	in := bufio.NewScanner(f)
	for in.Scan() {
		line := in.Text()
		match := re.FindStringSubmatch(line)

		year, err := strconv.Atoi(match[1])
		check(err)

		month, err := strconv.Atoi(match[2])
		check(err)

		day, err := strconv.Atoi(match[3])
		check(err)

		hour, err := strconv.Atoi(match[4])
		check(err)

		minute, err := strconv.Atoi(match[5])
		check(err)

		message := match[6]
		if message[0] == 'G' {
			guard, err := strconv.Atoi(match[7])
			check(err)
		}
	}
	check(in.Err())
}
