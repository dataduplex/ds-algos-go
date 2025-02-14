# Algos and DS


## Basics
### sets and strings
- Number of possible pairs in an array
  - N * (N-1)/2
  - Run time complexity to generate all pairs is O(N ^ 2)
- Number of possible subsets in an array.
  - 2 ^ N
  - Time complexity to generate all such sets O((2 ^ N) * N)
    - 2^N for the algorithm, and for every possibility, creating/copying to output makes it 2^N * N
  - Spacecomplexity to generate all such sets 
    - O((2 ^ N) * N) if output is included as incurred space.
    - O(N) if output is exluded and considering just the working variables.
- Number of possible substrings in an array.
  - N * (N + 1) / 2
  - Time and space complexities - same explanation as given to subsets.
- Number of possible k-size substrings in an array.
  - N - k + 1
  - Time and space complexities - same explanation as given to subsets.
- Number of possible k combinations in an array.
  - C(N, K) => N! / (K! * (N-K)!)
  - Time and space complexities - same explanation as given to subsets.
- Number of possible k permutations in an array.
  - P(N, K) => N! / (N - K)!
  - Time and space complexities - same explanation as given to subsets.

### sorting


### iteration and recursion

## Blind 75 Notes

### (128) Longest Consecutive Sequence
https://leetcode.com/problems/longest-consecutive-sequence/description/?envType=problem-list-v2&envId=oizxjoit

- First we can store all unique elements in a map or a set for fast checking if n+1 or n-1 is present or not.
- Then for each element we can count back as far as possible (checking each decrement is present in the map) and count front as far as possible (checking each increment is present in the map) to find the length of it's chain. If we do this for every element, we'll have a max chain length at the end.
- Clearly we are doing repeated work here. We can avoid that by just counting front (or just the increments). So if we encounter n, and n-1 is present in the map, we can skip counting and move on to the next element. The rationale here is, when we get to n-1 as part of main loop, our increment process will cover n.
- That's it. that's the main algorithm, of course at each step we compare against global max and update it accordingly.
- Intially I was getting stuck trying to come up with an O(N) algorithm at the outset. Sometimes O(N) solution will become clear from an O(N^2) or another sub-optimal solution.
- So if you're getting stuck, it's better to start with a sub-optimal solution and see if we can find a way to optimize it. 

### (3) Longest Substring Without Repeating Characters
https://leetcode.com/problems/longest-substring-without-repeating-characters/description/?envType=problem-list-v2&envId=oizxjoit

- This problem is similar to the previous one, but much easier to reason (We don't need to explore multiple directions)
- The intuition is that this a sliding window problem, and it is.
- We need two indices, start and end to track the sliding window. A map to keep track of elements in the window, so we can detect duplicate.
- We keep expanding the window till a duplicate was found. In this case, the end position remains same, but start position is adjusted to duplicate_first_occurence_index + 1. As a result of this all elements from original start_index to duplication_first_occurrence_index will be removed from the map.
- That's the algorithm, as usual at end step of expansion window, it is checked against the max and updated accordingly.
- It's clear this is an O(N) algorithm.



