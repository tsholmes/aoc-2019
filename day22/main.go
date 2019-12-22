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
	cardCount := 10007
	lines := strings.Split(input, "\n")

	cards := make([]int, cardCount)
	for i := range cards {
		cards[i] = i
	}

	for _, line := range lines {
		switch {
		case strings.HasPrefix(line, "deal with increment"):
			inc, _ := strconv.Atoi(line[len("deal with increment "):])
			ncards := make([]int, cardCount)
			for i := range cards {
				j := (inc * i) % cardCount
				ncards[j] = cards[i]
			}
			cards = ncards
		case line == "deal into new stack":
			for i := 0; i < len(cards)/2; i++ {
				j := len(cards) - i - 1
				cards[i], cards[j] = cards[j], cards[i]
			}
		default:
			cut, _ := strconv.Atoi(line[len("cut "):])
			ncards := make([]int, cardCount)
			for i := 0; i < len(cards); i++ {
				j := (i + cardCount - cut) % cardCount
				ncards[j] = cards[i]
			}
			cards = ncards
		}
	}
	for i := range cards {
		if cards[i] == 2019 {
			fmt.Println(i)
		}
	}
}

func part2() {
	cardCount := 119315717514047
	repCount := 101741582076661
	target := 2020
	// cardCount := 10007
	// repCount := 1
	// target := 2306
	lines := strings.Split(input, "\n")

	muls := [2]*big.Int{
		big.NewInt(1),
		big.NewInt(0),
	}
	for i := len(lines) - 1; i >= 0; i-- {
		line := lines[i]
		switch {
		case strings.HasPrefix(line, "deal with increment"):
			// i * inc = [a * inc^-1, b * inc^-1]
			inc, _ := strconv.Atoi(line[len("deal with increment "):])
			sub := big.NewInt(0).ModInverse(big.NewInt(int64(inc)), big.NewInt(int64(cardCount)))
			muls[0].Mul(muls[0], sub)
			muls[1].Mul(muls[1], sub)
		case line == "deal into new stack":
			// cardCount - i - 1 = [-a, -b-1]
			muls[0].Neg(muls[0])
			muls[1].Sub(big.NewInt(-1), muls[1])
		default:
			cut, _ := strconv.Atoi(line[len("cut "):])
			// i + cut = [a, b + cut]
			muls[1].Add(muls[1], big.NewInt(int64(cut)))
		}
	}

	mod := big.NewInt(int64(cardCount))
	muls[0].Mod(muls[0], mod)
	muls[1].Mod(muls[1], mod)

	// muls = [a, b]
	// f(i) = a*i + b MOD cardCount

	v := big.NewInt(int64(target))
	// a^n
	an := big.NewInt(0).Exp(muls[0], big.NewInt(int64(repCount)), mod)

	// san = sum(a^n for n in [0, repCount-1]) = a(a^repCount - 1) / (a - 1)
	san := big.NewInt(0).Add(an, big.NewInt(-1))
	denom := big.NewInt(0).Add(muls[0], big.NewInt(-1))
	denom = denom.ModInverse(denom, mod)
	san = san.Mul(san, denom)
	v2 := big.NewInt(0).Mul(san, muls[1])

	v = v.Mul(v, an)
	v = v.Add(v, v2)
	v = v.Mod(v, mod)

	res := v.Int64()
	if res < 0 {
		res += int64(cardCount)
	}
	fmt.Println(res)
}

const input = `deal with increment 21
deal into new stack
deal with increment 52
deal into new stack
deal with increment 19
deal into new stack
deal with increment 65
cut -8368
deal into new stack
cut -3820
deal with increment 47
cut -2965
deal with increment 38
deal into new stack
cut 9165
deal with increment 62
cut 3224
deal with increment 72
cut 4120
deal with increment 40
cut -9475
deal with increment 52
cut 4437
deal into new stack
cut 512
deal with increment 33
cut -9510
deal into new stack
cut -6874
deal with increment 56
cut -47
deal with increment 7
deal into new stack
deal with increment 13
cut 4530
deal with increment 67
deal into new stack
cut 8584
deal with increment 26
cut 4809
deal with increment 72
cut -8003
deal with increment 24
cut 1384
deal into new stack
deal with increment 13
deal into new stack
cut -1171
deal with increment 72
cut 6614
deal with increment 61
cut 1412
deal with increment 16
cut -4808
deal into new stack
deal with increment 51
cut -8507
deal with increment 60
cut 1202
deal with increment 2
cut -4030
deal with increment 4
cut -9935
deal with increment 57
cut -6717
deal with increment 5
deal into new stack
cut 3912
deal with increment 64
cut 5152
deal into new stack
deal with increment 68
deal into new stack
cut 49
deal with increment 31
cut 4942
deal with increment 44
cut -4444
deal with increment 47
cut -5533
deal with increment 51
cut -3045
deal with increment 67
cut -1711
deal with increment 46
cut -6247
deal into new stack
cut 969
deal with increment 55
cut 7549
deal with increment 62
cut -9083
deal with increment 54
deal into new stack
cut -3337
deal with increment 62
cut 1777
deal with increment 39
cut 3207
deal with increment 13`
