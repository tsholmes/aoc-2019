# [Day 7: Amplification Circuit](https://adventofcode.com/2019/day/7)

After this problem, I decided to refactor out the IntCode machine and move it top-level to `intcode.go`. My original solution for this problem is still here as `part1Orig` and `part2Orig` (with `part1` and `part2` the newer implementation using the factored out IntCode machine).

## Part 1

##### `00:09:20 #100`

This problem was the first that uses the existing IntCode machine without modifying it. Since the output of each IntCode machine was the input to the next, I just ran them one at a time, feeding the output of the previous as the input to the next. The solution here was just to try all permutations of 1 to 5 and keep track of the max value found. I was too lazy to write a proper permutation so I just did 5 nested for loops over 1 to 5 and skipped if it wasn't a permutation (my code to check if it wasn't a permutation is extremely suspect and bit me in part 2).

## Part 2

##### `0:23:25 #44`

This one was a bit more complicated than the first because you couldn't run the machines one at a time since the continuously fed into each other. You could probably do this by stopping once you hit an output instruction and running the next one so you didn't need true parallelism, but it was easier for me to just hook them up with channels and run them in separate goroutines. The rest of my solution was the same, just checking each permutation. For the last machine I manually fed it back to the input channel of the first, saving the last one I saw along the way. I lost 5 minutes or so because my only check for distinct permutations was that the product was equal to `1*2*3*4*5`, not realizing that `1*2*3*4*5 = 2*2*2*3*5`. I fixed this by adding in a check that the sum was also equal. I'm not 100% sure this is correct even for 1 through 5 (it's definitely not correct for larger ranges).
