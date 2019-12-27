# [Day 12: The N-Body Problem](https://adventofcode.com/2019/day/12)

## Part 1

##### `00:08:34 #41`

This part was just a simulation following the rules they gave. You loop through once updating the velocities, then loop through a second time updating the positions. After 1000 iterations of that, you just run the energy calculation provided.

## Part 2

##### `00:31:01 #56`

Part 2 was a lot more interesting since you couldn't reasonably simulate it long enough to actually see a duplicate state (although I implemented that real quick and ran it in the background while I though of something better). The first key observation for this one is realizing that each axis is entirely independent. The second key observation (which I can't prove, but feels intuitive) is that because the velocity change is always linear, the entire system is always cyclical. This means that the first duplicate state seen will always be returning everything to their initial positions with velocity zero. This means you don't actually need to keep track of and check against every state seen, you just need to run it until you return to the initial state. Then, since each axis is independent, you just need to run until you see each axis return to its initial state. Then you have separate cycle lengths for x, y, and z. These will line up for the first time at their least common multiple `lcm(cx, cy, cz)` which is just `(cx*cy*cz)/(gcf(cx,cy,cz)^2)`.
