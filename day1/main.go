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
	var total int64
	for _, line := range strings.Split(input, "\n") {
		x, _ := strconv.ParseInt(line, 10, 64)
		total += (x / 3) - 2
	}
	fmt.Println(total)
}

func part2() {
	var calc func(val int64) int64
	memo := map[int64]int64{}
	calc = func(val int64) int64 {
		if val <= 6 {
			return 0
		}
		res, ok := memo[val]
		if !ok {
			res = (val / 3) - 2 + calc((val/3)-2)
		}
		memo[val] = res
		return res
	}

	var total int64
	for _, line := range strings.Split(input, "\n") {
		x, _ := strconv.ParseInt(line, 10, 64)
		total += calc(x)
	}
	fmt.Println(total)
}

const input = `123265
68442
94896
94670
145483
93807
88703
139755
53652
52754
128052
81533
56602
96476
87674
102510
95735
69174
136331
51266
148009
72417
52577
86813
60803
149232
115843
138175
94723
85623
97925
141772
63662
107293
130779
147027
88003
77238
53184
149255
71921
139799
84851
104899
92290
74438
55631
58655
140496
110176
138718
104768
93177
53212
129572
69877
139944
116062
51362
135245
59682
128705
98105
69172
89244
109048
88690
62124
53981
71885
59216
107718
146343
138788
73588
51648
122227
54507
59283
101230
93080
123120
148248
102909
91199
105704
113956
120368
75020
103734
81791
87323
77278
123013
58901
136351
121295
132994
84039
76813`
