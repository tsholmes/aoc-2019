package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	type chem struct {
		count int
		name  string
	}
	type reac struct {
		in  []chem
		out chem
	}
	parsec := func(s string) chem {
		ps := strings.Split(s, " ")
		v, _ := strconv.ParseInt(ps[0], 10, 64)
		return chem{
			count: int(v),
			name:  ps[1],
		}
	}
	lines := strings.Split(input, "\n")
	rs := map[string]reac{}
	for _, line := range lines {
		ps := strings.Split(line, " => ")
		ins := strings.Split(ps[0], ", ")

		in := make([]chem, len(ins))
		for i, s := range ins {
			in[i] = parsec(s)
		}
		out := parsec(ps[1])
		rs[out.name] = reac{
			in:  in,
			out: out,
		}
	}

	depths := map[string]int{}
	var search func(string, int)
	search = func(name string, depth int) {
		v, ok := depths[name]
		if !ok || depth > v {
			depths[name] = depth
		}
		if name == "ORE" {
			return
		}
		r := rs[name]
		for _, in := range r.in {
			search(in.name, depth+1)
		}
	}

	search("FUEL", 0)

	var names []string
	for k := range rs {
		names = append(names, k)
	}
	sort.Slice(names, func(i, j int) bool {
		return depths[names[i]] < depths[names[j]]
	})

	counts := map[string]int{"FUEL": 1}
	for _, name := range names {
		r := rs[name]
		rcount := (counts[name] + r.out.count - 1) / r.out.count
		for _, in := range r.in {
			counts[in.name] += in.count * rcount
		}
	}

	fmt.Println(counts["ORE"])
}

func part2() {
	type chem struct {
		count int
		name  string
	}
	type reac struct {
		in  []chem
		out chem
	}
	parsec := func(s string) chem {
		ps := strings.Split(s, " ")
		v, _ := strconv.ParseInt(ps[0], 10, 64)
		return chem{
			count: int(v),
			name:  ps[1],
		}
	}
	lines := strings.Split(input, "\n")
	rs := map[string]reac{}
	for _, line := range lines {
		ps := strings.Split(line, " => ")
		ins := strings.Split(ps[0], ", ")

		in := make([]chem, len(ins))
		for i, s := range ins {
			in[i] = parsec(s)
		}
		out := parsec(ps[1])
		rs[out.name] = reac{
			in:  in,
			out: out,
		}
	}

	depths := map[string]int{}
	var search func(string, int)
	search = func(name string, depth int) {
		v, ok := depths[name]
		if !ok || depth > v {
			depths[name] = depth
		}
		if name == "ORE" {
			return
		}
		r := rs[name]
		for _, in := range r.in {
			search(in.name, depth+1)
		}
	}

	search("FUEL", 0)

	var names []string
	for k := range rs {
		names = append(names, k)
	}
	sort.Slice(names, func(i, j int) bool {
		return depths[names[i]] < depths[names[j]]
	})

	find := func(fuel int) int {
		counts := map[string]int{"FUEL": fuel}
		for _, name := range names {
			r := rs[name]
			rcount := (counts[name] + r.out.count - 1) / r.out.count
			for _, in := range r.in {
				counts[in.name] += in.count * rcount
			}
		}
		return counts["ORE"]
	}

	max := sort.Search(1000000000000, func(i int) bool {
		return find(i) >= 1000000000000
	})
	if find(max) > 1000000000000 {
		max--
	}
	fmt.Println(max, find(max))
}

const input = `4 JWXL => 8 SNBF
23 MPZQF, 10 TXVW, 8 JWXL => 6 DXLB
1 SNZDR, 5 XMWHC, 1 NJSC => 7 MHSB
2 TDHD, 11 TXVW => 4 RFNZ
2 VRCD, 1 FGZG, 3 JWXL, 1 HQTL, 2 MPZQF, 1 GTPJ, 5 HQNMK, 10 CQZQ => 9 QMTZB
3 SRDB, 2 ZMVLP => 3 DHFD
1 DFQGF => 1 CVXJR
193 ORE => 3 TRWXF
23 MFJMS, 4 HJXJH => 1 WVDF
5 TRWXF => 5 RXFJ
4 GZQH => 7 SNZDR
160 ORE => 4 PLPF
1 PLPF => 5 NJSC
2 QKPZ, 2 JBWFL => 7 HBSC
15 DXLB, 1 TDHD, 9 RFNZ => 5 DBRPW
7 PLPF, 4 GMZH => 7 PVNX
3 JWXL, 1 XWDNT, 4 CQZQ => 2 TPBXV
2 SNZDR => 9 WQWT
1 WMCF => 2 XWDNT
1 DFQGF, 8 FGZG => 5 LMHJQ
168 ORE => 9 GMZH
18 PVNX, 3 RXFJ => 4 JBWFL
5 WQWT => 1 CQZQ
6 QMTZB, 28 NVWM, 8 LMHJQ, 1 SNBF, 15 PLPF, 3 KMXPQ, 43 WVDF, 52 SVNS => 1 FUEL
164 ORE => 9 RXRMQ
2 MFJMS, 1 HJXJH, 7 WVDF => 7 NXWC
8 QDGBV, 1 WMCF, 2 MHSB => 6 HQTL
1 XMWHC => 8 MLSK
2 GMZH, 1 RXRMQ => 2 GZQH
4 MPZQF, 7 WVDF => 9 KHJMV
4 ZMVLP, 19 MLSK, 1 GZQH => 8 MFJMS
1 HQTL, 1 SXKQ => 2 PWBKR
3 SXKQ, 16 TXVW, 4 SVNS => 5 PSRF
4 MPZQF, 3 SVNS => 9 QDGBV
7 NXWC => 8 FGZG
7 TDHD, 1 WQWT, 1 HBSC => 9 TXVW
14 JBWFL => 5 LMXB
1 VRCD, 3 KHJMV => 3 RTBL
16 DHFD, 2 LBNK => 9 SXKQ
1 QDGBV, 1 NJSC => 6 JWXL
4 KHJMV => 3 HQNMK
5 GZQH => 6 LBNK
12 KHJMV, 19 FGZG, 3 XWDNT => 4 VRCD
5 DHFD, 3 MLSK => 8 QKPZ
4 KHJMV, 1 CQDR, 3 DBRPW, 2 CQZQ, 1 TPBXV, 15 TXVW, 2 TKSLM => 5 NVWM
2 KHJMV => 5 CQDR
1 CVXJR => 8 SVNS
35 RXFJ, 5 NJSC, 22 PVNX => 9 HJXJH
5 LMXB => 3 DFQGF
1 RXFJ => 2 SRDB
20 TPBXV, 1 RTBL, 13 PWBKR, 6 RFNZ, 1 LMXB, 2 CVXJR, 3 PSRF, 25 MPZQF => 9 KMXPQ
1 MHSB, 8 MPZQF => 3 TDHD
6 DHFD, 3 LBNK => 7 WMCF
1 SRDB => 7 ZMVLP
3 RXFJ => 8 XMWHC
1 MPZQF => 8 TKSLM
9 JBWFL, 22 WQWT => 8 MPZQF
12 HBSC, 15 TKSLM => 1 GTPJ`
