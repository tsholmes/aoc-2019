# [Day 21: Springdroid Adventure](https://adventofcode.com/2019/day/21)

## Part 1

##### `00:16:26 #105`

This problem was just an exercise in writing assembly code with only 2 temporary registers (technically the `J` register was a temporary until the program was done). The logic of my program was just if any of the next 3 spots (`A`, `B`, or `C`) was a hole, and the tile 4 spots away (`D` where you land if you jump immediately) wasn't a hole, then jump.

```
NOT A T
OR T J
// J = !A
NOT B T
OR T J
// J = !A || !B
NOT C T
OR T J
// J = !A || !B || !C
AND D J
// J = (!A || !B || !C) && D
WALK
```

I could have cut out the second instruction by just making the first `NOT` write to `J`.

## Part 2

##### `00:19:26 #18`

Part 2 just used a harder sequence of holes that required you too look ahead further, up to 9 spots (which new read-only registers `E` through `I`). The logic of my program was to start with my solution for part 1 (if a hole in `A` `B` or `C` and ground at `D`), and add in the condition that either there must be ground on the tile after I land (`E`) or if I needed to jump again immediately, the tile I landed on after that (`H`) was ground. This avoids the case where you jump sooner than needed (because there is a hole at `C`) only to find that your only choice is to jump into a 3 wide pit. There are likely cases that could be possible to solve with only 15 instructions that this doesn't cover, but it worked for me.

```
NOT A T
OR T J
NOT B T
OR T J
NOT C T
OR T J
AND D J
// J = (!A || !B || !C) && D
NOT J T
AND J T
// the two above instructions set T to false so I can OR into it
OR E T
OR H T
AND T J
// J = (!A || !B || !C) && D && (E || H)
RUN
```
