package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	ps := strings.Split(input, "")
	ds := make([]int, len(ps))
	for i, p := range ps {
		ds[i], _ = strconv.Atoi(p)
	}

	base := []int{0, 1, 0, -1}

	next := make([]int, len(ds))
	for phase := 0; phase < 100; phase++ {
		for i := range ds {
			next[i] = 0
			for j, d := range ds {
				p := ((j + 1) / (i + 1)) % len(base)
				next[i] += base[p] * d
			}
			next[i] = next[i] % 10
			if next[i] < 0 {
				next[i] = -next[i]
			}
		}
		ds, next = next, ds
	}
	fmt.Println(ds[:8])
}

func part2() {
	ps := strings.Split(input, "")
	ds := make([]int, len(ps))
	for i, p := range ps {
		ds[i], _ = strconv.Atoi(p)
	}
	ds2 := make([]int, len(ps)*10000)
	for i := range ds2 {
		ds2[i] = ds[i%len(ds)]
	}
	ds = ds2

	offset := 0
	for _, v := range ds[:7] {
		offset *= 10
		offset += v
	}
	ds2 = ds2[offset:]

	cs := make([]int, len(ds2))
	for i := range cs {
		binom := big.NewInt(0).Binomial(99+int64(i), 99)
		binom = big.NewInt(0).Mod(binom, big.NewInt(10))
		cs[i] = int(binom.Int64())
	}

	for i := 0; i < 8; i++ {
		for j := 1; i+j < len(ds2); j++ {
			ds2[i] += ds2[i+j] * cs[j]
			ds2[i] = ds2[i] % 10
		}
	}
	fmt.Println(ds2[:8])
}

const input2 = `12345678`
const input = `59768092839927758565191298625215106371890118051426250855924764194411528004718709886402903435569627982485301921649240820059827161024631612290005106304724846680415690183371469037418126383450370741078684974598662642956794012825271487329243583117537873565332166744128845006806878717955946534158837370451935919790469815143341599820016469368684893122766857261426799636559525003877090579845725676481276977781270627558901433501565337409716858949203430181103278194428546385063911239478804717744977998841434061688000383456176494210691861957243370245170223862304663932874454624234226361642678259020094801774825694423060700312504286475305674864442250709029812379`
