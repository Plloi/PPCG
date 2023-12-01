package main

import (
	"regexp"
	"unicode"
)

var firstNumber = regexp.MustCompile("(zero|one|two|three|four|five|six|seven|eight|nine|[0-9])")
var lastNumber = regexp.MustCompile("([0-9]|enin|thgie|neves|xis|evif|ruof|eerht|owt|eno|orez)")

var numberMap = map[string]int{
	"zero":  0,
	"orez":  0,
	"0":     0,
	"one":   1,
	"eno":   1,
	"1":     1,
	"two":   2,
	"owt":   2,
	"2":     2,
	"three": 3,
	"eerht": 3,
	"3":     3,
	"four":  4,
	"ruof":  4,
	"4":     4,
	"five":  5,
	"evif":  5,
	"5":     5,
	"six":   6,
	"xis":   6,
	"6":     6,
	"seven": 7,
	"neves": 7,
	"7":     7,
	"eight": 8,
	"thgie": 8,
	"8":     8,
	"nine":  9,
	"enin":  9,
	"9":     9,
}

func day1Part1(calibrate []string) int {
	total := 0
	for _, line := range calibrate {
		total += 10*getFirstDigit(line) + getLastDigit(line)
	}
	return total
}

func day1Part2(calibrate []string) int {
	total := 0
	for _, line := range calibrate {
		total += 10*getFirstNumber(line) + getLastNumber(line)
	}
	return total
}

func getFirstDigit(line string) int {
	for _, v := range line {
		if unicode.IsDigit(v) {
			return int(v - '0')
		}
	}
	return 0
}

func getLastDigit(line string) int {
	for i := len(line) - 1; i >= 0; i-- {
		v := rune(line[i])
		if unicode.IsDigit(v) {
			return int(v - '0')
		}
	}
	return 0
}

func getFirstNumber(line string) int {
	res := firstNumber.FindString(line)
	return numberMap[res]
}

func getLastNumber(line string) int {
	res := lastNumber.FindString(reverse(line))
	return numberMap[res]

}
