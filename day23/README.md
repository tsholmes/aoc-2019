# Day 23: Category Six

## Part 1

##### `00:11:05 #85`

Part 1 just involved running 50 of the provided IntCode program and passing the outputs of each machine as inputs to the correct machine. One small change here compared to other IntCode problems was that reading was non-blocking, so you had to make it default to `-1` if there was nothing to be read. You likely could have done these in lock-step or some other way that wasn't actually concurrent, but I just spun them up each in their own goroutine.

## Part 2

##### `00:19:29 #82`

The (somewhat) difficult part of the second half was knowing when to send the NAT packet to machine 0, mostly for me since everything was running in different goroutines and I had to lock around reads/writes. I did this by having each machine keep a flag of whether the last read it did defaulted to `-1`, and when machine 0 went to read, if it and everything else had its last read as a `-1` and the queues for every machine were empty, read the NAT packet instead of a `-1`. Then you just needed to keep a map of NAT packets that you had delivered to machine 0 and run the simulation until you saw a duplicate. Since my solution ran each machine in a separate goroutine, it wasted a LOT of cpu cycles just looping over `-1` inputs repeatedly, causing it to take upwards of 20 seconds until I got a duplicate NAT packet. This would likely be an order of magnitude faster if I instead suspended the machines after reading a `-1` until something sent a packet to it.
