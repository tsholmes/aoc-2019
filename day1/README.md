# [Day 1: The Tyranny of the Rocket Equation](https://adventofcode.com/2019/day/1)

I forgot and missed this one :(

## Part 1

##### `16:52:11 #24231`

Just a simple equation. Loop over all the numbers and sum up the values.

## Part 2

##### `16:55:18 #21599`

Similar to part 1, but you have to do the equation recursively and sum up the values until the cost of what is left is `<= 0` (`mass <= 6`). I did this with memoization to avoid repeated calculations, but after realized that the calculation is `log(n)`, so you can definitely just do it with a simple loop.
