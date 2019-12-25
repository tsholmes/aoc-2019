# Day 8: Space Image Format

## Part 1

##### `00:03:21 #27`

This one was fairly simple, just splitting the input into layers, and counting the number of zeros and ones. Then just keep an running minimum of zeros, and the product of the counts for the current minimum.

## Part 2

##### `00:09:56 #85`

For this one, I just went through each pixel and checked the layers to find the final value of it. I lost a few minutes because I was comparing against the number `2` instead of the character `'2'`, and while doing that switched it to look bottom up instead of top-down, but top-down was definitely the easier move. Then I replaced everything that wasn't a `1` with a space and printed it out. I lost another minute because first I had it transposed, and then `1` characters don't really fill a lot of space, so I spent a bit squinting at the output trying to read the letters. Using `#` or something more space filling would have made it a lot easer
