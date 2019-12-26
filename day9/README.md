# Day 9: Sensor Boost (IntCode Part 3: The Last One)

## Part 1

##### `00:27:27 #436`

This is the last problem that modifies the IntCode machine. First you need to support relative mode for parameters, and to implement the instruction to adjust the relative base. Then you have to add support for infinite memory. For this problem I did that by adding a map that stores all the values outside of the initial array, and then on every read/write checking if the address was outside the initial range (I later just made the entire memory a map). The last requirement was to support "large numbers", which turned out to just be 64 bit integers, which I was already using. I lost a lot of time because in this refactoring I accidentally swapped the address and the value in the writes, which produces a lot of confusing output.

## Part 2

##### `00:28:06 #410`

Part 2 wasn't actually anything additional. You just change the value to the single input.
