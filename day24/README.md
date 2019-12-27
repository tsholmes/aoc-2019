# [Day 24: Planet of Discord](https://adventofcode.com/2019/day/24)

## Part 1

##### `00:05:55 #12`

This problem is a simple cellular automaton similar to Conway's game of life, but instead of 8 neighbors with 2,3 liveness and 3 genesis, you have 4 neighbors with 1 liveness and 1,2 genesis. Its done on a small finite grid where edges have fewer neighbors, so its fairly straightforward to implement by just calculating the next state based on the current state, keeping track of previously seen states until you see a duplicate state, then running the formula for score against that duplicate state.

## Part 2

##### `00:18:17 #4`

Part 2 at first seems rather complicated to implement, since you may have a continuously growing state to manage. The observation I used to simplify my implementation is that for simulating 200 steps, you can't create bugs more than 200 levels in each direction from the starting state (it actually takes 2 steps to cross a layer, so you only really need 100 levels in each direction). I did this by having the state consist of 401 levels, with the initial state being the state of the middle level (layer 200 0-indexed). Then you just need to determine the neighbors of each cell and run the simulation 200 steps forward. I just did this by hardcoding the one extra neighbor you get when going out of the current grid for each direction, and hardcoding the 5 neighbors you get when going into the middle grid for each direction (where layer 0 has no outer neighbors and layer 400 has no inner neighbors). Then the overall implementation was fairly similar to part 1, with the addition of skipping the center square for each layer, and looping over all 401 layers to calculate the next state.
