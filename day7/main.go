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
	run := func(input []int64) int64 {
		opstr := strings.Split(INPUT, ",")
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

		val := func(p int64, fullOp int64, off int64) int64 {
			mode := (fullOp / off) % 10
			if mode == 0 {
				return ops[p]
			} else {
				return p
			}
		}

		for {
			fullOp := next()
			op := fullOp % 100
			var p1, p2, p3 int64
			switch op {
			case 1:
				p1 = next()
				p2 = next()
				p3 = next()
				ops[p3] = val(p1, fullOp, 100) + val(p2, fullOp, 1000)
			case 2:
				p1 = next()
				p2 = next()
				p3 = next()
				ops[p3] = val(p1, fullOp, 100) * val(p2, fullOp, 1000)
			case 3:
				p1 = next()
				p2, input = input[0], input[1:]
				ops[p1] = p2
			case 4:
				p1 = next()
				return val(p1, fullOp, 100)
			case 5:
				p1 = next()
				p2 = next()
				if val(p1, fullOp, 100) != 0 {
					pos = val(p2, fullOp, 1000)
				}
			case 6:
				p1 = next()
				p2 = next()
				if val(p1, fullOp, 100) == 0 {
					pos = val(p2, fullOp, 1000)
				}
			case 7:
				p1 = next()
				p2 = next()
				p3 = next()
				if val(p1, fullOp, 100) < val(p2, fullOp, 1000) {
					ops[p3] = 1
				} else {
					ops[p3] = 0
				}
			case 8:
				p1 = next()
				p2 = next()
				p3 = next()
				if val(p1, fullOp, 100) == val(p2, fullOp, 1000) {
					ops[p3] = 1
				} else {
					ops[p3] = 0
				}
			case 99:
				panic("FAIL")
			default:
				panic(op)
			}
		}
	}

	max := int64(0)

	for v1 := 1; v1 <= 5; v1++ {
		for v2 := 1; v2 <= 5; v2++ {
			for v3 := 1; v3 <= 5; v3++ {
				for v4 := 1; v4 <= 5; v4++ {
					for v5 := 1; v5 <= 5; v5++ {
						if v1*v2*v3*v4*v5 != 1*2*3*4*5 {
							continue
						}

						output := int64(0)
						for _, v := range []int{v1, v2, v3, v4, v5} {
							output = run([]int64{int64(v - 1), output})
						}
						if output > max {
							max = output
						}
					}
				}
			}
		}
	}
	fmt.Println(max)
}

func part2() {
	run := func(input chan int64, output chan int64) {
		opstr := strings.Split(INPUT, ",")
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

		val := func(p int64, fullOp int64, off int64) int64 {
			mode := (fullOp / off) % 10
			if mode == 0 {
				return ops[p]
			} else {
				return p
			}
		}

		for {
			fullOp := next()
			op := fullOp % 100
			var p1, p2, p3 int64
			switch op {
			case 1:
				p1 = next()
				p2 = next()
				p3 = next()
				ops[p3] = val(p1, fullOp, 100) + val(p2, fullOp, 1000)
			case 2:
				p1 = next()
				p2 = next()
				p3 = next()
				ops[p3] = val(p1, fullOp, 100) * val(p2, fullOp, 1000)
			case 3:
				p1 = next()
				p2 = <-input
				ops[p1] = p2
			case 4:
				p1 = next()
				output <- val(p1, fullOp, 100)
			case 5:
				p1 = next()
				p2 = next()
				if val(p1, fullOp, 100) != 0 {
					pos = val(p2, fullOp, 1000)
				}
			case 6:
				p1 = next()
				p2 = next()
				if val(p1, fullOp, 100) == 0 {
					pos = val(p2, fullOp, 1000)
				}
			case 7:
				p1 = next()
				p2 = next()
				p3 = next()
				if val(p1, fullOp, 100) < val(p2, fullOp, 1000) {
					ops[p3] = 1
				} else {
					ops[p3] = 0
				}
			case 8:
				p1 = next()
				p2 = next()
				p3 = next()
				if val(p1, fullOp, 100) == val(p2, fullOp, 1000) {
					ops[p3] = 1
				} else {
					ops[p3] = 0
				}
			case 99:
				close(output)
				return
			default:
				panic(op)
			}
		}
	}

	max := int64(0)

	for v1 := 1; v1 <= 5; v1++ {
		for v2 := 1; v2 <= 5; v2++ {
			for v3 := 1; v3 <= 5; v3++ {
				for v4 := 1; v4 <= 5; v4++ {
					for v5 := 1; v5 <= 5; v5++ {
						if v1*v2*v3*v4*v5 != 1*2*3*4*5 || v1+v2+v3+v4+v5 != 1+2+3+4+5 {
							continue
						}

						initIn := make(chan int64, 200)
						in := initIn
						var out chan int64
						for i, v := range []int{v1, v2, v3, v4, v5} {
							in <- int64(v + 4)
							if i == 0 {
								in <- 0
							}
							out = make(chan int64, 200)
							go run(in, out)
							in = out
						}

						output := int64(0)
					outer:
						for {
							select {
							case v, ok := <-out:
								if !ok {
									break outer
								}
								output = v
								initIn <- v
							}
						}
						if output > max {
							fmt.Println(v1, v2, v3, v4, v5)
							max = output
						}
					}
				}
			}
		}
	}
	fmt.Println(max)
}

const INPUT = `3,8,1001,8,10,8,105,1,0,0,21,34,43,64,85,98,179,260,341,422,99999,3,9,1001,9,3,9,102,3,9,9,4,9,99,3,9,102,5,9,9,4,9,99,3,9,1001,9,2,9,1002,9,4,9,1001,9,3,9,1002,9,4,9,4,9,99,3,9,1001,9,3,9,102,3,9,9,101,4,9,9,102,3,9,9,4,9,99,3,9,101,2,9,9,1002,9,3,9,4,9,99,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,1,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,2,9,4,9,99,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,99,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,101,1,9,9,4,9,99,3,9,101,1,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,99,3,9,101,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,102,2,9,9,4,9,99`
