package aoc

import (
	"fmt"
	"strconv"
	"strings"
)

func LoadIntCode(src string, input func() int64, output func(int64)) IntCode {
	opstr := strings.Split(src, ",")
	ops := make([]int64, len(opstr))
	for i, s := range opstr {
		ops[i], _ = strconv.ParseInt(s, 10, 64)
	}

	return IntCode{
		Mem:    ops,
		Input:  input,
		Output: output,
	}
}

func FixedInput(input ...int64) func() int64 {
	return func() int64 {
		var v int64
		v, input = input[0], input[1:]
		return v
	}
}

func PrintOutput() func(int64) {
	return func(v int64) {
		fmt.Println("OUTPUT", v)
	}
}

func SingleOutput(out *int64) func(int64) {
	return func(v int64) {
		*out = v
	}
}

func ChanInput(ch <-chan int64) func() int64 {
	return func() int64 {
		return <-ch
	}
}

func ChanOutput(ch chan<- int64) func(int64) {
	return func(v int64) {
		ch <- v
	}
}

type IntCode struct {
	Mem    []int64
	Pos    int64
	Input  func() int64
	Output func(int64)
	Done   bool

	modes [2]int64
	modep int
}

func (i *IntCode) RunToEnd() {
	for !i.Done {
		i.Step()
	}
}

func (i *IntCode) Step() {
	if i.Done {
		panic("done")
	}
	op := i.nextOp()
	var p1, p2, p3 int64
	switch op {
	case 1:
		p1 = i.nextVal()
		p2 = i.nextVal()
		p3 = i.nextPtr()
		i.Mem[p3] = p1 + p2
	case 2:
		p1 = i.nextVal()
		p2 = i.nextVal()
		p3 = i.nextPtr()
		i.Mem[p3] = p1 * p2
	case 3:
		p1 = i.nextPtr()
		p2 = i.Input()
		i.Mem[p1] = p2
	case 4:
		p1 = i.nextVal()
		i.Output(p1)
	case 5:
		p1 = i.nextVal()
		p2 = i.nextVal()
		if p1 != 0 {
			i.Pos = p2
		}
	case 6:
		p1 = i.nextVal()
		p2 = i.nextVal()
		if p1 == 0 {
			i.Pos = p2
		}
	case 7:
		p1 = i.nextVal()
		p2 = i.nextVal()
		p3 = i.nextPtr()
		if p1 < p2 {
			i.Mem[p3] = 1
		} else {
			i.Mem[p3] = 0
		}
	case 8:
		p1 = i.nextVal()
		p2 = i.nextVal()
		p3 = i.nextPtr()
		if p1 == p2 {
			i.Mem[p3] = 1
		} else {
			i.Mem[p3] = 2
		}
	case 99:
		i.Done = true
	default:
		panic(op)
	}
}

func (i *IntCode) next() int64 {
	v := i.Mem[i.Pos]
	i.Pos++
	return v
}

func (i *IntCode) nextOp() int64 {
	fullOp := i.next()
	op := fullOp % 100
	i.modes[0] = (fullOp / 100) % 10
	i.modes[1] = (fullOp / 1000) % 10
	i.modep = 0
	return op
}

func (i *IntCode) nextVal() int64 {
	v := i.next()
	mode := i.modes[i.modep]
	i.modep++

	switch mode {
	case 0:
		return i.Mem[v]
	case 1:
		return v
	default:
		panic(mode)
	}
}

func (i *IntCode) nextPtr() int64 {
	i.modep++
	return i.next()
}
