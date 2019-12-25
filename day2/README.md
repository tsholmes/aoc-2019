# Day 2: 1202 Program Alarm (IntCode: Genesis)

Missed my alarm and started late on this one :(

# Part 1

##### `00:47:41 #3239`

Fairly simple implementation of this one. There's no jumping or anything, so you can just run the instructions in order until you hit a 99.

# Part 2

##### `00:50:49 #2406`

No changes needed to the IntCode implementation, but you do have to run it multiple times now. Since the answer you had to give was `100 * noun + verb`, I assumed both were in the range `[0, 99]` and turned out to be right. All it took was two for loops around the solution to part 1.
