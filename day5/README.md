# [Day 5: Sunny with a Chance of Asteroids (IntCode Part 2)](https://adventofcode.com/2019/day/5)

## Part 1

##### `00:10:03 #21`

This adds I/O to the IntCode machine, and immediate mode for parameters. I just copied my IntCode implementation from day 2 and added the new instructions. For now input is just a slice where one element is consumed each time the input instruction is called. It will panic if too many are read. Output just prints to standard out.

## Part 2

##### `00:13:12 #10`

This is where things start to heat up (at least in the IntCode world). Here we need to add 2 comparison instructions (set less than, set equal), and two jump instructions (jump if zero, jump if nonzero). Now you can't just loop through the program in order once, you have to keep a pointer to the next instruction and move it around. The problem description talks about making sure not to move the pointer after the instruction when jumping. I just moved the pointer as I was reading the opcode and instructions, so I was never modifying it after the instruction and didn't have to special case any behavior for jumping.
