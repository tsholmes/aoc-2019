package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func main() {
	part1()
	part2()
}

func gcf(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcf(b, a%b)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func part1() {
	lines := strings.Split(input, "\n")

	var as [][2]int
	for i, line := range lines {
		for j := 0; j < len(line); j++ {
			if line[j] == '#' {
				as = append(as, [2]int{j, i})
			}
		}
	}

	max := 0
	var maxa [2]int
	for i, a := range as {
		count := 0
		for j, b := range as {
			if j == i {
				continue
			}

			dx, dy := b[0]-a[0], b[1]-a[1]

			var ox, oy int
			if dx == 0 {
				oy = dy / abs(dy)
			} else if dy == 0 {
				ox = dx / abs(dx)
			} else {
				f := gcf(abs(dx), abs(dy))
				ox, oy = dx/f, dy/f
			}

			x, y := ox, oy

			v := true
			for x != dx || y != dy {
				if lines[y+a[1]][x+a[0]] == '#' {
					v = false
				}
				x += ox
				y += oy
			}
			if v {
				count++
			}
		}
		if count > max {
			max = count
			maxa = a
		}
	}
	fmt.Println(max, maxa)
}

func part2() {
	lines := strings.Split(input, "\n")

	var as [][2]int
	for i, line := range lines {
		for j := 0; j < len(line); j++ {
			if line[j] == '#' {
				as = append(as, [2]int{j, i})
			}
		}
	}

	a := [2]int{26, 28}
	// a := [2]int{11, 13}

	vap := make(map[int][][2]int)

	for _, b := range as {
		if a == b {
			continue
		}
		dx, dy := b[0]-a[0], b[1]-a[1]

		var ox, oy int
		if dx == 0 {
			oy = dy / abs(dy)
		} else if dy == 0 {
			ox = dx / abs(dx)
		} else {
			f := gcf(abs(dx), abs(dy))
			ox, oy = dx/f, dy/f
		}

		x, y := ox, oy

		pass := 0
		for x != dx || y != dy {
			if lines[y+a[1]][x+a[0]] == '#' {
				pass++
			}
			x += ox
			y += oy
		}
		vap[pass] = append(vap[pass], b)
	}

	off := 0
	var p [][2]int
	for i := 0; ; i++ {
		if off+len(vap[i]) < 200 {
			off += len(vap[i])
		} else {
			p = vap[i]
			break
		}
	}

	angle := func(i int) float64 {
		dx := p[i][0] - a[0]
		dy := p[i][1] - a[1]
		a := math.Atan2(float64(dy), float64(dx))
		if a < -math.Pi/2.0 {
			a += 2.0 * math.Pi
		}
		return a
	}

	sort.Slice(p, func(i, j int) bool {
		return angle(i) < angle(j)
	})

	fmt.Println(p[199-off])
}

const input2 = `.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##`

const input = `#....#.....#...#.#.....#.#..#....#
#..#..##...#......#.....#..###.#.#
#......#.#.#.....##....#.#.....#..
..#.#...#.......#.##..#...........
.##..#...##......##.#.#...........
.....#.#..##...#..##.....#...#.##.
....#.##.##.#....###.#........####
..#....#..####........##.........#
..#...#......#.#..#..#.#.##......#
.............#.#....##.......#...#
.#.#..##.#.#.#.#.......#.....#....
.....##.###..#.....#.#..###.....##
.....#...#.#.#......#.#....##.....
##.#.....#...#....#...#..#....#.#.
..#.............###.#.##....#.#...
..##.#.........#.##.####.........#
##.#...###....#..#...###..##..#..#
.........#.#.....#........#.......
#.......#..#.#.#..##.....#.#.....#
..#....#....#.#.##......#..#.###..
......##.##.##...#...##.#...###...
.#.....#...#........#....#.###....
.#.#.#..#............#..........#.
..##.....#....#....##..#.#.......#
..##.....#.#......................
.#..#...#....#.#.....#.........#..
........#.............#.#.........
#...#.#......#.##....#...#.#.#...#
.#.....#.#.....#.....#.#.##......#
..##....#.....#.....#....#.##..#..
#..###.#.#....#......#...#........
..#......#..#....##...#.#.#...#..#
.#.##.#.#.....#..#..#........##...
....#...##.##.##......#..#..##....`
