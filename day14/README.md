# [Day 14: Space Stoichiometry](https://adventofcode.com/2019/day/14)

## Part 1

##### `00:16:08 #24`

The main problem to solve here was finding the order to run the reactions in so that for a specific product, you always knew exactly how much you needed. You only want to consider each reaction once and run it a specific number of times. This avoids discarding excess products multiple times which could add up to wasting more than a full reaction of product. To find this order, you treat the reactions as a graph with the products having nodes to their reagents. Since each chemical other than ORE is produced by exactly one reaction, you have a directed acyclic graph. You then topologically sort the graph by by distance from the FUEL node and run the reactions in reverse. This means you start with 1 FUEL and then go through all the possible reactions in reverse, tracking how many of that reactions inputs you need to produce the number of outputs you expect. You do this by just dividing the amount the reaction produces by the amount you expect and rounding up. From one of the examples:

```
9 ORE => 2 A (distance 2)
8 ORE => 3 B (distance 2)
7 ORE => 5 C (distance 2)
3 A, 4 B => 1 AB (distance 1)
5 B, 7 C => 1 BC (distance 1)
4 C, 1 A => 1 CA (distance 1)
2 AB, 3 BC, 4 CA => 1 FUEL (distance 0)
```

This gives a possible ordering of `FUEL, CA, BC, AB, C, B, A`. Other orderings are valid as long as reactions with a higher distance are later in the list. I'm omitting the products from the reactions we processed backwards, because you won't see them again, so they are irrelevant. Then running over the reactions you get:

```
State: 1 FUEL

2 AB, 3 BC, 4 CA => 1 FUEL (run once, ceil(1/1) = 1)
State: 2 AB, 3 BC, 4 CA

4 C, 1 A => 1 CA (run 4 times, ceil(4/1) = 4)
State: 2 AB, 3 BC, 16 C, 4 A

5 B, 7 C => 1 BC (run 3 times, ceil(3/1) = 3)
State: 2 AB, 4 A, 15 B, 37 C

3 A, 4 B => 1 AB (run twice, ceil(2/1) = 2)
State: 10 A, 23 B, 37 C

7 ORE => 5 C (run 8 times, ceil(37/5) = 8)
State: 10 A, 23 B, 56 ORE

8 ORE => 3 B (run 8 times, ceil(23/3) = 8)
State: 10 A, 120 ORE

9 ORE => 2 A (run 5 times, ceil(10/2) = 5)
State: 165 ORE
```

## Part 2

##### `00:18:29 #13`

This part involves finding out how many FUEL you can produce from 1 trillion ORE. Since we already wrote an algorithm to figure out how many ORE we need for 1 fuel, we can re-use that. By updating the initial state from 1 FUEL to N FUEL, we can determine the minimum ore needed to produce N FUEL (this will likely be different than just multiplying the ORE for 1 FUEL by N, since this ignores all the discarded products). Then we can binary search over a range of N values (I chose 1 to 1 trillion to be safe), and found the lowest number of fuel that took > 1 trillion ORE, subtracting 1 to get the max.
