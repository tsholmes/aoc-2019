# Day 18: Many-Worlds Interpretation

## Part 1

##### `00:28:35 #49`

This problem involved pathfinding through a maze while the geometry of the maze changed slightly by picking up the keys. This meant that you couldn't just pre-compute all the path lengths, but had to update it along the way (there is some things you can precompute though that I will discuss later). My solution here was to recursively search over all orders of picking up the keys. At each state I would go through each key that I hadn't already picked up, find the shortest path (via BFS) to that key that didn't cross any doors I hadn't unlocked yet, or any keys that I hadn't picked up. This solution only finished in a somewhat reasonable amount of time (~15s) due to aggressive memoization and short circuiting. I memoized the length of the shortest path from each point to another given a set of picked up keys, and kept track of the minimum number of steps taken to reach a specific state of position and picked up keys, returning early if I was already at the same or higher number of steps as a previous pass.

My solution could get a LOT faster with a few improvements. First, the maze can be reduced to a graph containing nodes for the starting position and each key and door. The edges in that graph are the minimum distances between the two points WITHOUT crossing any other node. This gives a much smaller space that you can pathfind over much quicker, while still maintaining correctness (if you aren't taking the shortest path without crossing a key/door, then you are going to that key/door and taking the shortest path from it on). Since this would add edges with different weights, you would probably want to switch the search algorithm to Djikstra's instead of a simple BFS. The second speedup would come from changing how I decide which keys to go for next. Instead of trying to pathfind to each key, taking the ones that have a valid path (which requires N searches for N remaining keys), I could just do a single path search, returning a list of the reachable keys and their path lengths.

## Part 2

##### `00:40:29 #5`

Part 2 was just a slight modification of part 1 where you had 4 independent robots in 4 separate mazes that weren't connected to each other. You could only move one robot at the time, so the solution was almost exactly the same as part 1, you just need to keep track of 4 current locations instead of 1. I sped this part up by precomputing which keys were in each robot's maze, so when I was checking the shortest path to a specific key, I only needed to pathfind to it from one position (instead of 4). Since the search space on this problem was much smaller, it ends up running an order of magnitude faster than part 1.

This problem would become a bit more difficult and interesting (but still tractable) if the robots were allowed to move simultaneously. This would force you to keep track of how many steps it took to reach each key, and ensuring you don't cross that door until the key is gathered.
