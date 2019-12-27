# Day 22: Slam Shuffle

## Part 1

##### `00:09:42 #55`

Part 1 dealt with a fairly small number of cards (10007) going through ~100 shuffle steps, so I just shuffled the whole deck and then found where card 2019 was at the end. There is a solution that doesn't require actually shuffling though. Since you are looking for the index a specific card ends up at, you can just track that one card forwards through the shuffle steps:

```
deal with increment:
f(i, inc) = i * inc MOD cardCount

deal into new stack (reverse):
f(i) = cardCount - i - 1
OR
f(i) = -i - 1 MOD cardCount

cut:
f(i, cut) = i - cut MOD cardCount
```

One interesting observation here is that the `deal with increment` step only works for increments that are coprime with the number of cards. The card count here is prime, so any non-zero increment is valid.

## Part 2

##### `01:34:26 #72`

Part 2 is my favorite problem of this year by far (even though it took me over an hour and a half to get right). This involved finding out which card ended up at a specific index where you start with an extremely large number of cards (~100 trillion) and repeat the shuffling process an extremely large number of times (also ~100 trillion). Either of these two limitations alone make it so that shuffling the whole deck or running through all the steps are both far too costly operations to actually perform. The card count is still prime, so any non-zero increment for `deal with increment` is still valid. Since you are looking for which card ends up at a specific index, you take that index and go through the shuffle steps backwards to find the starting index of the card that ends up there.

Since the card count and the number of steps are so large, you need to find a closed form equation to calculate the result of the shuffle steps. I started this by finding a closed form for running through the shuffle steps backwards a single time, getting a formula of the form `f'(i) = a*i + b MOD cardCount`.
With no shuffle steps, this starts as `f'(i) = i`, or `a, b = 1, 0`. Then going through each shuffle step in reverse, you transform your `a, b` pair based on the step being performed, which involves doing the inverse of the forward steps:

Since we are dealing with numbers `MOD cardCount` (a ring in mathematical terms), division is only defined for divisors that have a modular inverse: numbers that are coprime with the modulus (which we already determined are all non-zero numbers since the modulus is prime). Division is then transformed to multiplying by the modular inverse:

```
deal with increment:
f(i, inc) = i * inc MOD cardCount
f'(i, inc) = i / inc MOD cardCount
f'(i, inc) = i * modinv(inc) MOD cardCount
```

Reversing is a symmetric operation, so the inverse is the same as the forward shuffle:

```
deal into new stack (reverse):
f(i) = -i - 1 MOD cardCount
f'(i) = -i - 1 MOD cardCount
```

Inverting the cut operation just reverses the sign of the cut:

```
cut:
f(i, cut) = i - cut MOD cardCount
f'(i, cut) = i + cut MOD cardCount
```

Looking at these shuffle operations in terms of a linear equation of `a*i+b` gives you the new coefficients `a, b`:

```
deal with increment:
f'(a*i+b, inc) = (a*i+b) * modinv(inc)                     MOD cardCount
               = (a * modinv(inc)) * i + (b * modinv(inc)) MOD cardCount
giving new coefficients
f'([a, b], inc) = [a * modinv(inc), b * modinv(inc)]

deal into new stack (reverse):
f'(a*i+b) = -(a*i+b) - 1      MOD cardCount
          = (-a)*i + (-1 - b) MOD cardCount
giving new coefficients
f'([a, b]) = [-a, -1 - b]

cut:
f'(a*i+b) = (a*i+b) + cut   MOD cardCount
          = a*i + (b + cut) MOD cardCount
giving new coefficients
f'([a, b]) = [a, b + cut]
```

After running through all the shuffle steps backwards once, you end up with a single linear equation `f(i) = a*i + b` that performs all the shuffle steps backwards once. The next step is to find a closed form equation for running the shuffle steps a certain number of times (recursively):

```
f_1(i) = f(i) = a*i + b

f_2(i) = f(f(i)) = a*(a*i + b) + b
                 = (a^2)*i + a*b + b
                 = (a^2)*i + (a+1)*b MOD cardCount

f_3(i) = f(f(f(i))) = a*(a*(a*i +b) + b) + b
                    = (a^3)*i + (a^2)*b + a*b + b
                    = (a^3)*i + (a^2 + a + 1)*b MOD cardCount

Which generalizes to:
f_N(i) = (a^N)i + (a^(N-1) + a^(N-2) + ... + a^1 + a^0)*b MOD cardCount
```

The closed form for that summation (which I googled) is:

```
a^(N-1) + ... + a^0 = (a^N - 1) / (a - 1)
                    = (a^N - 1) * modinv(a - 1)
```

This gives us final coefficients of:

```
f_N([a, b]) = [a^N, (a^N - 1) * modinv(a - 1) * b]
```

The final result is then just calculating `f_reps(i)` with the `a` and `b` coefficients you got from running through the shuffle steps backwards once and the repetition count given in the problem description. I used the golang `big.Int` class for calculating the modular inverse and for doing exponentiation with a modulus (but the inputs sizes here should stay within 64-bit uints as long as you mod by the cardcount after every operation).
