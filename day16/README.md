# Day 16: Flawed Frequency Transmission

## Part 1

##### `00:10:17 #59`

The input size for part 1 was relatively small, so you can just implement the rules provided and calculate the result after 100 steps.

## Part 2

##### `01:48:15 #382`

This was a significant jump in difficulty from previous problems, and it took me quite a while to even realize a possible solution. The key realization here was that the offset that was the first 7 digits of the input was more than halfway through the total input (after duplicating 10k times). This means that you only ever got through half of the repeating pattern, and it gave you some nice properties that allow a closed-form solution. For example, for 10 digits, if we just look at the coefficients for the last 5 digits, we get:

```
#|5 6 7 8 9 (In Digit)
-+---------
5|1 1 1 1 1
6|0 1 1 1 1
7|0 0 1 1 1
8|0 0 0 1 1
9|0 0 0 0 1
(Out digit)
```

This means the next value of each digit is just the sum of itself and all the digits after it. After running through some examples manually to find the coefficients after N iterations (and looking up some sequences in OEIS), I found that for each output digit, the coefficient for a digit M indices to the right of it after N iterations was `binomial(M+N-1, N-1)`, which means after 100 iterations it was `binomial(M+99, 99)`. Then the solution is just to calculate the binomial coefficients mod 10 (up to the total number of digits after the repeating 10k times and cutting off everything before the offset), and then calculating the value of the first 8 digits.
