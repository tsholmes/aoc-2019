package main

import (
	"fmt"
	"strconv"
)

func main() {
	part1()
	part2()
}

func part1() {
	count := 0
	for x := inputMin; x <= inputMax; x++ {
		str := strconv.FormatInt(int64(x), 10)
		lastDigit := byte(0)
		same := false
		valid := true
		for _, d := range []byte(str) {
			if d == lastDigit {
				same = true
			}
			if d < lastDigit {
				valid = false
			}
			lastDigit = d
		}
		if valid && same {
			count++
		}
	}
	fmt.Println(count)
}

func part2() {
	count := 0
	for x := inputMin; x <= inputMax; x++ {
		str := strconv.FormatInt(int64(x), 10)
		lastDigit := byte(0)
		sameCount := 1
		valid := true
		counts := []int{}
		for _, d := range []byte(str) {
			if d == lastDigit {
				sameCount++
			} else {
				counts = append(counts, sameCount)
				sameCount = 1
			}
			if d < lastDigit {
				valid = false
			}
			lastDigit = d
		}
		counts = append(counts, sameCount)
		sameValid := false
		for _, c := range counts {
			if c == 2 {
				sameValid = true
			}
		}
		if valid && sameValid {
			count++
		}
	}
	fmt.Println(count)
}

const inputMin = 125730
const inputMax = 579381
