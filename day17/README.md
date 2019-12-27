# [Day 17: Set and Forget](https://adventofcode.com/2019/day/17)

## Part 1

##### `00:10:17 #206`

Part 1 just involved running the IntCode program and parsing its output looking for paths that crossed. I assumed none crossed on the edge and just found tiles which were marked as a path and had all of its 4 neighbors marked as path (or the current position of the robot).

## Part 2

##### `00:39:45 #62`

Part 2 made you find a path that covered every position that wasn't a hole. This would be trivial if not for the restrictions placed on the length of the commands given to the robot to follow the path. I started this one by generating a series of "turn" and "move N spaces" commands that followed the entire path. I assumed that if you could go straight, you should always go straight (as this reduced the total number of turn commands needed, and we are length restricted). From there, you could write an algorithm that tried all possible subprograms of valid length until you found one that reduced the final command list to 20 characters or less. This algorithm didn't immediately come to me, so I just looked at the long-list of commands and manually picked out segments that were repeated (which is why my solution just has the resulting command list hard coded).
