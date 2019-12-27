# Day 25: Cryostasis

## Part 1

##### `00:47:51 #221`

This one would have been a pain to algorithmically search the space and solve (given that there were traps that quit early and a trap that caused an infinite loop). I instead played the game manually by hooking up stdin/stdout to the IntCode machine and just wrote each command I did down. If I accidentally hit a trap, I just restarted and copied in the commands I used to get to the previous state. Once I had collected all 7 items (it would be 8 but after some manual testing early on I found the `dark matter` alone made me too heavy so I ignored it), I wrote a quick program to drop each subset of items and try going through the checkpoint until one of them didn't give me the error message that I was ejected.

This problem took me a lot longer than it should have because I initially though there were only 5 items, and I tried all combinations manually (after dropping the dark matter). Only then did I realize that I must be missing some items and had to explore more until I found the rest.

## Part 2

##### `00:47:54 #169`

This one is a freebie once you've solved all the other problems.
