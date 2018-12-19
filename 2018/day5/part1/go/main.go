package main

import (
	"fmt"
	"io/ioutil"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func React(a byte, b byte) bool {
	return a != b && unicode.ToLower(rune(a)) == unicode.ToLower(rune(b))
}

func main() {
	s, err := ioutil.ReadFile("../input")
	check(err)
	s = s[:len(s)-1]

	i := 0
	length := len(s)
	for j := 1; j < len(s); j++ {
		if React(s[i], s[j]) {
			length -= 2
			s[i] = ' '
			s[j] = ' '

			for i > 0 && s[i] == ' ' {
				i--
			}

			if i < 0 {
				j++
				i = j
			}
		} else {
			i = j
		}
	}

	fmt.Println(length)
}
