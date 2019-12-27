# [Day 20: Donut Maze](https://adventofcode.com/2019/day/20)

## Part 1

##### `00:17:11 #31`

The most difficult part of this problem was parsing the input. I ended up doing this by checking each tile in the maze that was a `.`, then checking two characters in each direction to see if they were both letters, which would mark it as a portal entrance/exit (or the start/end). The ordering of the two letters to form the name was dependent on which direction you looked (for up and left, the letter 2 tiles away came first, where for down and right, the letter 1 tile away came first). Parsing the input was further confounded by the fact that the input lines were not all the same length, so you had to check wether an index was valid per line instead of just using a fixed width and height. After finding the portals (and keeping a map of portal entrance to exit) and start/end positions, this became a simple graph search, which I did with a BFS.

## Part 2

##### `00:22:56 #4`

Part 2 was similar to part 1 except you also had to keep track of the current level as well as the x,y coordinates, and you only finished if you hit ZZ on level 0. The outer portals were only valid if you weren't on level 0. I distinguished inner from outer portals by a simple heuristic: if the portal entrance is more than 5 tiles from the outer edge of the input, it was in inner portal. After this, it was again just a BFS.
