# Day 11: Space Police

## Part 1

##### `00:09:17 #55`

This was the first problem that involved running an IntCode machine while continuously reading output and using it to decide the next input. I just did this by making it do I/O over channels and running it in a separate goroutine. The rest of part 1 was fairly simple, just keeping a map of squares that were colored and their current color, and keeping track of the current position based on the direction it said it turned. The result is then the number of distinct squares it ever wrote to. I used a map to keep track of the squares, so the result is just the length of the map.

## Part 2

##### `00:12:23 #47`

The only change here is just making `0,0` start as a white square instead of black. Then I found the min/max x and y coordinates written to in the map, and drew out the entire grid covering those squares to the console. I originally had it transposed which made it very difficult to read, but after swapping the x and y coordinates the letters became clear.
