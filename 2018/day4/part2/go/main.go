package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type timestamp struct {
	year   int
	month  int
	day    int
	hour   int
	minute int
}

type event struct {
	awake bool
	guard int
}

type chrono []timestamp

func (t chrono) Len() int {
	return len(t)
}

func (t chrono) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t chrono) Less(i, j int) bool {
	if t[i].year == t[j].year {
		if t[i].month == t[j].month {
			if t[i].day == t[j].day {
				if t[i].hour == t[j].hour {
					return t[i].minute < t[j].minute
				}
				return t[i].hour < t[j].hour
			}
			return t[i].day < t[j].day
		}
		return t[i].month < t[j].month
	}
	return t[i].year < t[j].year
}

func main() {
	history := make(map[timestamp]event)
	times := make([]timestamp, 0, 10)
	sleepiness := make(map[int]int)
	freq := make(map[int][60]int)

	re := regexp.MustCompile(`^\[(\d+)\-(\d+)\-(\d+) (\d+):(\d+)\] (Guard #(\d+) begins shift|falls asleep|wakes up)$`)

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

		guard := -1
		message := match[6]
		if message[0] == 'G' {
			guard, err = strconv.Atoi(match[7])
			check(err)
		}

		awake := message[0] != 'f'

		t := timestamp{year, month, day, hour, minute}
		e := event{awake, guard}

		history[t] = e
		times = append(times, t)
	}
	check(in.Err())

	sort.Sort(chrono(times))
	guard := -1
	bedtime := -1
	bestGuard := -1
	bestSleepiness := -1
	for _, t := range times {
		ev := history[t]

		if ev.guard != -1 {
			guard = ev.guard
		} else if ev.awake && bedtime != -1 {
			fs := freq[guard]
			bestFreq := -1
			for minute := bedtime; minute < t.minute; minute++ {
				fs[minute]++

				if bestFreq < fs[minute] {
					bestFreq = fs[minute]
				}
			}
			freq[guard] = fs

			if sleepiness[guard] < bestFreq {
				sleepiness[guard] = bestFreq
				if bestSleepiness < bestFreq {
					bestSleepiness = bestFreq
					bestGuard = guard
				}
			}

			bedtime = -1
		} else if !ev.awake && bedtime == -1 {
			bedtime = t.minute
		}
	}

	bestMinute := -1
	bestFreq := -1
	for minute, f := range freq[bestGuard] {
		if bestFreq < f {
			bestFreq = f
			bestMinute = minute
		}
	}

	fmt.Println(bestGuard * bestMinute)
}
