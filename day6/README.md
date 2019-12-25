# Day 6: Universal Orbit Map

## Part 1

##### `00:07:21 #171`

Since the orbit list is a tree, this is just the sum of the depths of each node in the tree. I memoized the depths while searching, but the number of nodes was fairly small (< 1000) so I probably didn't need to.

## Part 2

##### `00:11:04 #81`

I did this problem as a graph search, adding forward and backward links for each pair in the input. This worked fine (after a few bug fixes), but there is a much easier solution. This problem is effectively the least common ancestor problem for two nodes in a tree. You find the least common ancestor, and the solution is `depth(YOU)-1 + depth(SAN)-1 - 2*LCA(YOU, SAN)`.
