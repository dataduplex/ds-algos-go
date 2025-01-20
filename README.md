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




## another title


