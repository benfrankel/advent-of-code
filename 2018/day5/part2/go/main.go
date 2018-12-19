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

func Equiv(a byte, b byte) bool {
	return unicode.ToLower(rune(a)) == unicode.ToLower(rune(b))
}

func React(a byte, b byte) bool {
	return a != b && Equiv(a, b)
}

func Reduce(s []byte, a byte) int {
	length := len(s)

	if len(s) <= 0 {
		return 0
	}

	var i int
	for i = 0; i < len(s) && Equiv(s[i], a); i++ {
		length--
		s[i] = ' '
	}

	if i >= len(s) {
		return 0
	}

	for j := i + 1; j < len(s); j++ {
		for ; j < len(s) && Equiv(s[j], a); j++ {
			length--
			s[j] = ' '
		}

		if j >= len(s) {
			break
		}

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

	return length
}

func main() {
	s, err := ioutil.ReadFile("../input")
	check(err)
	s = s[:len(s)-1]

	length := Reduce(s, '!')

	t := make([]byte, length)
	i := 0
	for _, c := range s {
		if c != ' ' {
			t[i] = c
			i++
		}
	}

	bestLength := length
	for unit := byte('a'); unit <= byte('z'); unit++ {
		tmp := make([]byte, len(t))
		copy(tmp, t)
		reducedLength := Reduce(tmp, unit)

		if bestLength > reducedLength {
			bestLength = reducedLength
		}
	}

	fmt.Println(bestLength)
}
