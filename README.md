# leetcode_climbing_stairs

This repository is for solve [climbing-stairs](https://leetcode.com/problems/climbing-stairs/)

## problem description

You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

## Examples

### Example1
```
Input: n = 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps
```
### Example2
```
Input: n = 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step
```
## Observation
```
Notice that

There exists a relation

climb n stairs methods = first Step is One Step + first Step is Two Steps

                       = climb n - 1 stairs method + climb n - 2 stairs method

for n >= 3
                        
```
## implemtations

There are two implemention for solve this problem

Solution 1:

Create an array to store previous to accumlate which take Space Complexity O(n)

and iterate n step to compute n times

Thus, Time Complexity is O(n)

Solution 2:

Use iteration to accumlate previous two steps result

iterate n steps to compute 

Time Complexity is O(n)