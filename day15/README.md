# [Day 15: Oxygen System](https://adventofcode.com/2019/day/15)

## Part 1

##### `00:33:31 #155`

This was another interactive IntCode problem. Part 1 involved sending commands to explore an unknown space for a specific object. I did this by just doing a search every time it asked for a movement command to path-find to the nearest unknown tile until I found the target spot. I quit immediately on finding it and returned the length of the shortest path from the origin to the target that only goes over known squares. This had a chance of being very wrong of the shape of the space was complicated enough, and to get a provably correct answer, you needed to full explore the space before returning the path length. I ended up getting lucky and the answer I got after a partial search was correct. I lost a fair bit of time on this one because I had the direction commands going clockwise (north, east, south, west) instead of opposite pairs (north, south, east, west), so I thought i was going infinitely north-east when really I was just moving back and forth between the same two spots.

The runtime of my solution here is much slower than it needs to be. I am re-doing the search for every movement command instead of just saving a path and doing a new search when my path runs out. I'm also doing two separate searches to find a new tile, and then find a path to that tile. This could be combined into a single search where the state contains the first move made on the path to each tile.

## Part 2

##### `00:36:55 #89`

Part 2 was framed as a simulation problem of letting the gas expand one tile at a time, but in reality all you were searching for was the maximum shortest distance of any tile to the target position. This did require fully exploring the space. Once explored (there were no unknown tiles that had a valid path to them), you just breadth-first-search out from the target position, and return the maximum distance that you saw any tile for the first time.
