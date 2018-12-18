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

type claim struct {
	id int
	l  int
	r  int
	t  int
	b  int
}

func Intersects(a claim, b claim) bool {
	return a.id != b.id && a.l < b.r && b.l < a.r && a.t < b.b && b.t < a.b
}

func main() {
	intact := make(map[claim]bool)

	f, err := os.Open("../input")
	check(err)
	defer f.Close()

	in := bufio.NewScanner(f)
	re := regexp.MustCompile("^#(\\d+) @ (\\d+),(\\d+): (\\d+)x(\\d+)$")
	for in.Scan() {
		line := in.Text()
		match := re.FindStringSubmatch(line)

		id, err := strconv.Atoi(match[1])
		check(err)

		l, err := strconv.Atoi(match[2])
		check(err)

		t, err := strconv.Atoi(match[3])
		check(err)

		w, err := strconv.Atoi(match[4])
		check(err)

		h, err := strconv.Atoi(match[5])
		check(err)

		c1 := claim{id, l, l + w, t, t + h}

		intact[c1] = true
		for c2, _ := range intact {
			if Intersects(c1, c2) {
				intact[c1] = false
				intact[c2] = false
			}
		}
	}
	check(in.Err())

	for c, i := range intact {
		if i {
			fmt.Println(c.id)
			break
		}
	}
}
