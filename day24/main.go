package main

import (
	"fmt"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	var board [5][5]bool
	for i, line := range strings.Split(input, "\n") {
		for j := 0; j < 5; j++ {
			board[i][j] = line[j] == '#'
		}
	}

	seen := map[[5][5]bool]bool{}
	for {
		if seen[board] {
			break
		}
		seen[board] = true

		var nextBoard [5][5]bool
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				count := 0
				for _, dir := range [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
					i2, j2 := i+dir[0], j+dir[1]
					if i2 < 0 || i2 >= 5 || j2 < 0 || j2 >= 5 {
						continue
					}
					if board[i2][j2] {
						count++
					}
				}
				if board[i][j] {
					nextBoard[i][j] = count == 1
				} else {
					nextBoard[i][j] = count == 1 || count == 2
				}
			}
		}
		board = nextBoard
	}

	score := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			bit := 1 << uint((i*5)+j)
			if board[i][j] {
				score += bit
			}
		}
	}
	fmt.Println(score)
}

func part2() {
	var boards [401][5][5]bool
	for i, line := range strings.Split(input, "\n") {
		for j := 0; j < 5; j++ {
			boards[200][i][j] = line[j] == '#'
		}
	}

	neighbors := func(level int, i int, j int) [][3]int {
		var res [][3]int
		dirs := [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
		ins := [][][2]int{
			{{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 4}},
			{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}},
			{{4, 0}, {4, 1}, {4, 2}, {4, 3}, {4, 4}},
			{{0, 4}, {1, 4}, {2, 4}, {3, 4}, {4, 4}},
		}
		outs := [][2]int{
			{3, 2},
			{2, 3},
			{1, 2},
			{2, 1},
		}
		for di, dir := range dirs {
			i2, j2 := i+dir[0], j+dir[1]
			if i2 == 2 && j2 == 2 {
				if level == 400 {
					continue
				}
				for _, in := range ins[di] {
					res = append(res, [3]int{level + 1, in[0], in[1]})
				}
			} else if i2 < 0 || i2 >= 5 || j2 < 0 || j2 >= 5 {
				if level == 0 {
					continue
				}
				res = append(res, [3]int{level - 1, outs[di][0], outs[di][1]})
			} else {
				res = append(res, [3]int{level, i2, j2})
			}
		}
		return res
	}

	for iter := 0; iter < 200; iter++ {
		var nextBoard [401][5][5]bool
		for level := 0; level < 401; level++ {
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					if i == 2 && j == 2 {
						continue
					}
					count := 0
					for _, n := range neighbors(level, i, j) {
						if boards[n[0]][n[1]][n[2]] {
							count++
						}
					}
					if boards[level][i][j] {
						nextBoard[level][i][j] = count == 1
					} else {
						nextBoard[level][i][j] = count == 1 || count == 2
					}
				}
			}
		}
		boards = nextBoard
	}

	score := 0
	for _, b := range boards {
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				if b[i][j] {
					score++
				}
			}
		}
	}
	fmt.Println(score)
}

const input2 = `....#
#..#.
#..##
..#..
#....`

const input = `#####
.#.#.
.#..#
....#
..###`
