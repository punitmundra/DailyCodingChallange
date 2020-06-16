package main

/*
There's a staircase with N steps, and you can climb 1 or 2 steps at a time. 
Given N, write a function that returns the number of unique ways you can climb the staircase. The order of the steps matters.

For example, if N is 4, then there are 5 unique ways:

1, 1, 1, 1
2, 1, 1
1, 2, 1
1, 1, 2
2, 2
=================

n=1 1
n=2 2
n=3 3
n=4 5
n=5 8

https://afteracademy.com/blog/climbing-stairs-problem
*/

import (
	"fmt"
)

// CountTheWaysToClimbTheStairCase ...
func CountTheWaysToClimbTheStairCase(steps []int, climbStair int) {
	dp := make([]int, climbStair+1)
	dp[0] = 1
	for i := 1; i < (climbStair + 1); i++ {

		for _, val := range steps {
			diff := i - val
			if diff < 0 {
				break
			} else {
				dp[i] += dp[diff]
			}
		}
	}

	for i, val := range dp {
		fmt.Println("for steps :", i, " total ways :", val)
	}
}
func main() {

	steps := []int{2, 4, 6}
	stair := 5
	fmt.Println("steps:", steps)
	CountTheWaysToClimbTheStairCase(steps, stair)

}
