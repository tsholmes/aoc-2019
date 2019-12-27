# [Day 13: Care Package (Breakout!)](https://adventofcode.com/2019/day/13)

## Part 1

##### `00:03:15 #85`

The first part was extremely simple, just running the IntCode program and counting how many tiles are blocks at the end.

## Part 2

##### `00:14:12 #17`

Part 2 involved actually playing the game until you broke the last block. I took a gamble that a naive bot playing the game would eventually finish. Every time it asked for a move I would move the paddle left if the ball was the the left, and right if it was to the right. This took a while for it to finish, but it did finish and didn't require me to search for an optimal move by implementing the game rules myself.
