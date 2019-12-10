package aoc

import (
	"strconv"
	"strings"
)

func CommaNums(input string) []int64 {
	ss := strings.Split(input, ",")
	ns := make([]int64, len(ss))
	for i, s := range ss {
		ns[i], _ = strconv.ParseInt(s, 10, 64)
	}
	return ns
}

func CommaLinesNums(input string) [][]int64 {
	lines := strings.Split(input, "\n")
	ns := make([][]int64, len(lines))
	for i, line := range lines {
		ns[i] = CommaNums(line)
	}
	return ns
}

func SepLinesStrs(input string, sep string) [][]string {
	lines := strings.Split(input, "\n")
	ss := make([][]string, len(lines))
	for i, line := range lines {
		ss[i] = strings.Split(line, sep)
	}
	return ss
}

func CommaLinesStrs(input string) [][]string {
	return SepLinesStrs(input, ",")
}

func Digits(input string) []int {
	ds := make([]int, len(input))
	for i := 0; i < len(input); i++ {
		ds[i] = int(input[i] - '0')
	}
	return ds
}

func DigitLines(input string) [][]int {
	lines := strings.Split(input, "\n")
	ds := make([][]int, len(lines))
	for i, line := range lines {
		ds[i] = Digits(line)
	}
	return ds
}
