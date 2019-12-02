package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	opstr := strings.Split(input, ",")
	ops := make([]int64, len(opstr))
	for i, s := range opstr {
		ops[i], _ = strconv.ParseInt(s, 10, 64)
	}

	var pos int64

	next := func() int64 {
		val := ops[pos]
		pos++
		return val
	}

	ops[1] = 12
	ops[2] = 2

outer:
	for {
		op := next()
		var p1, p2, p3 int64
		switch op {
		case 1:
			p1 = next()
			p2 = next()
			p3 = next()
			ops[p3] = ops[p1] + ops[p2]
		case 2:
			p1 = next()
			p2 = next()
			p3 = next()
			ops[p3] = ops[p1] * ops[p2]
		case 99:
			break outer
		default:
			panic(op)
		}
	}
	fmt.Println(ops[0])
}

func part2() {
	var expected int64 = 19690720

	opstr := strings.Split(input, ",")
	initops := make([]int64, len(opstr))
	for i, s := range opstr {
		initops[i], _ = strconv.ParseInt(s, 10, 64)
	}

	for noun := int64(0); noun < 100; noun++ {
		for verb := int64(0); verb < 100; verb++ {
			ops := make([]int64, len(initops))
			copy(ops, initops)

			ops[1] = noun
			ops[2] = verb

			var pos int64

			next := func() int64 {
				val := ops[pos]
				pos++
				return val
			}

		outer:
			for {
				op := next()
				var p1, p2, p3 int64
				switch op {
				case 1:
					p1 = next()
					p2 = next()
					p3 = next()
					ops[p3] = ops[p1] + ops[p2]
				case 2:
					p1 = next()
					p2 = next()
					p3 = next()
					ops[p3] = ops[p1] * ops[p2]
				case 99:
					break outer
				default:
					panic(op)
				}
			}

			if ops[0] == expected {
				fmt.Println(noun, verb)
				return
			}
		}
	}

}

const input = `1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,6,1,19,1,19,10,23,2,13,23,27,1,5,27,31,2,6,31,35,1,6,35,39,2,39,9,43,1,5,43,47,1,13,47,51,1,10,51,55,2,55,10,59,2,10,59,63,1,9,63,67,2,67,13,71,1,71,6,75,2,6,75,79,1,5,79,83,2,83,9,87,1,6,87,91,2,91,6,95,1,95,6,99,2,99,13,103,1,6,103,107,1,2,107,111,1,111,9,0,99,2,14,0,0`
