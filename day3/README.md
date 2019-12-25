# Day 3: Crossed Wires

## Part 1

##### `00:06:29 #26`

Looking at the input I saw that the lengths of the segments were relatively small (around 1000 max), so instead of doing anything complicated, I could just keep track of all the points I crossed. My solution was to walk the instructions in the input in order, and whenever I hit a point I haven't seen before, keep track of the minimum manhattan distance from the origin.

## Part 2

##### `00:08:08 #7`

Very similar to part 1, but instead of looking at the manhattan distance, you store the number of steps taken so far for each point, and when you find a point you've seen before, keep track of the minimum steps taken for the already seen point.
