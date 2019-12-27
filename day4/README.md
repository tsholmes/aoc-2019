# [Day 4: Secure Container](https://adventofcode.com/2019/day/4)

## Part 1

##### `00:02:40 #61`

The input numbers were relatively small (6 digits), so we don't need to do anything fancy, we can just try all of them. For each number you convert it to digits (I just converted them to strings) and check that the digits are non-descending and there is at least one adjacent duplicate pair.

## Part 2

##### `00:06:29 #104`

A slightly more difficult set of rules. The same outer approach works (check every number in the range), but you have to make sure there is at least one adjacent duplicate pair that is not part of a larger run. There are simpler ways to do it than what I did (just check the digit before and after the pair), but my solution was to keep track of all the run lengths of digits, and at the end check if there were any that were exactly 2, while also keeping the check from part 1 that the digits are non-descending.
