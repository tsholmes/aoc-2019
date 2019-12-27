# [Day 19: Tractor Beam](https://adventofcode.com/2019/day/19)

## Part 1

##### `00:02:32 #32`

This problem is fairly straight-forward. Just run the IntCode program for each coordinate from (0, 0) to (49, 49) and sum up the outputs.

## Part 2

##### `00:09:52 #3`

For part 2, I made the assumption that the bottom-left corner of the 100x100 square would be against the left edge of the tractor beam. It feels like an intuitively correct assumption, but might not always be the case (it worked out for me though). My solution was to go down one row at a time, keeping track of the left edge, by moving the left edge of the previous row over until I found a spot that was affected by the tractor beam. Then I just checked the 4 corners of the square assuming that left edge was the bottom left of it (top-left is (x, y-99), bottom-right is (x+99, y), top-right is (x+99, y-99)). When all of the other corners were affected by the tractor beam, the top-left of the square (x, y-99) was the solution.
