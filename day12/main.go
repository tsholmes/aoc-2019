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

func sign(v int64) int64 {
	if v < 0 {
		return -1
	} else if v > 0 {
		return 1
	}
	return 0
}

func abs(v int64) int64 {
	if v < 0 {
		return -v
	}
	return v
}

func part1() {
	lines := strings.Split(input, "\n")
	poss := make([][3]int64, len(lines))
	vels := make([][3]int64, len(lines))
	for i, line := range lines {
		line = line[:len(line)-1]
		ps := strings.Split(line, ",")
		for j := 0; j < 3; j++ {
			poss[i][j], _ = strconv.ParseInt(ps[j][3:], 10, 64)
		}
	}

	for step := 0; step < 1000; step++ {
		for i := 0; i < len(poss); i++ {
			for j := 0; j < len(poss); j++ {
				if i == j {
					continue
				}
				pi, pj := poss[i], poss[j]
				d1, d2, d3 := pj[0]-pi[0], pj[1]-pi[1], pj[2]-pi[2]
				vels[i][0] += sign(d1)
				vels[i][1] += sign(d2)
				vels[i][2] += sign(d3)
			}
		}
		for i := range poss {
			poss[i][0] += vels[i][0]
			poss[i][1] += vels[i][1]
			poss[i][2] += vels[i][2]
		}
	}

	var e int64
	for i := range poss {
		p := poss[i]
		v := vels[i]
		e += (abs(p[0]) + abs(p[1]) + abs(p[2])) * (abs(v[0]) + abs(v[1]) + abs(v[2]))
	}
	fmt.Println(e)
}

func gcf(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcf(b, a%b)
}

func part2() {
	lines := strings.Split(input, "\n")
	poss := [4][3]int64{}
	vels := [4][3]int64{}
	for i, line := range lines {
		line = line[:len(line)-1]
		ps := strings.Split(line, ",")
		for j := 0; j < 3; j++ {
			poss[i][j], _ = strconv.ParseInt(ps[j][3:], 10, 64)
		}
	}

	initP := poss

	found := [3]bool{}
	cycles := [3]int{}

	for step := 1; !found[0] || !found[1] || !found[2]; step++ {
		for i := 0; i < len(poss); i++ {
			for j := 0; j < len(poss); j++ {
				if i == j {
					continue
				}
				pi, pj := poss[i], poss[j]
				d1, d2, d3 := pj[0]-pi[0], pj[1]-pi[1], pj[2]-pi[2]
				vels[i][0] += sign(d1)
				vels[i][1] += sign(d2)
				vels[i][2] += sign(d3)
			}
		}
		for i := range poss {
			poss[i][0] += vels[i][0]
			poss[i][1] += vels[i][1]
			poss[i][2] += vels[i][2]
		}

		for i := 0; i < 3; i++ {
			if !found[i] && [4]int64{vels[0][i], vels[1][i], vels[2][i], vels[3][i]} == [4]int64{} &&
				[4]int64{poss[0][i], poss[1][i], poss[2][i], poss[3][i]} ==
					[4]int64{initP[0][i], initP[1][i], initP[2][i], initP[3][i]} {
				found[i] = true
				cycles[i] = step
			}
		}
	}
	mul := gcf(gcf(cycles[0], cycles[1]), cycles[2])
	fmt.Println(cycles[0] * (cycles[1] / mul) * (cycles[2] / mul))
}

const input = `<x=3, y=-6, z=6>
<x=10, y=7, z=-9>
<x=-3, y=-7, z=9>
<x=-8, y=0, z=4>`
