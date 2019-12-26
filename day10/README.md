# Day 10: Monitoring Station

## Part 1

##### `00:13:20 #86`

This was the first of the math-based problems, a sign of things to come. The asteroids are essentially point masses, so they only block each-other if there is something exactly between them. I solved this one by checking each asteroid as a base station, and then going through each other asteroid to find its direction vector (`pBase - pOtherAsteroid`). In order to find asteroids that blocked each-other, you have to reduce the direction vector to the minimal integer offsets that point in the same direction. If the x or y difference between the base and the asteroid is zero, the other coordinate's difference is reduced to 1 or -1 depending if the original difference was positive or negative. For all other cases, you divide both coordinates by the greatest common factor of their absolute values `gcf(abs(dx), abs(dy))`. Then the solution is just the maximum number of distinct direction vectors seen with each asteroid as the base.

## Part 2

##### `00:33:13 #48`

For this problem you had to first determine which pass you hit the 200th asteroid, then find which asteroid in that pass it was (I didn't realize by looking at the answer to part 1 that for me it was the first pass, so I could have saved myself a bit of work). To do this, I went through the asteroids again, splitting them into passes. To determine which pass an asteroid was in, I found its minimum offset vector from the base (same as part 1), and then walked all multiples of that offset outwards from the base, counting how many asteroids I crossed. Then once I had them in passes, I picked which pass to use based on how many asteroids would be destroyed up to that pass, until it was >= 200. Then for that pass, I sorted the asteroids by their angle from the base (using `math.Atan2`). This gives an angle in the range `[-pi, pi]` with right as 0. I adjusted this range to `[-pi/2, 3*pi/2]` by just adding `2*pi` if the angle was `< -pi/2`. This put upwards as the minimum angle and sorted them clockwise. Then you just take the asteroid at index 199 minus the length of the previous passes (none in my case).
