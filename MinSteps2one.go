/*

Given a positive integer N, find the smallest number of steps it will take to reach 1.

There are two kinds of permitted 

Rule 1:
You may decrement N to N - 1.

Rule 2 : 
If a * b = N, you may decrement N to the larger of a and b.


For example, given 100, you can reach 1 in five steps with the following route: 

100 (R2: 10 * 10)-> 10(R1)  -> 9(Rule2: 3* 3) -> 3(R1) -> 2(R1) -> 1.

Answer : 5 

we applied total rules 5 times 


Below solution is for the BFS
*/

package main

import "fmt"

// Excluding 1 and n from the list of factor
func Factors(n int) []int {
	factors := []int{} 
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}

type Element struct {
	value int
	count int
}

func GetMinSteps(n int) int {
	queue := make([]Element, 0)
	visited := make(map[int]int)
	queue = append(queue, Element{value: n, count: 0})
	iterations := 0
	for len(queue) > 0 {
		element := queue[0]
		iterations++
		queue = queue[1:]
		if element.value == 1 {
			fmt.Println("iterations:", iterations, " For Number:", n, " Required totals steps:", element.count)
			return element.count
		}
		// Apply the Rule 1:
		if _, ok := visited[element.value-1]; ok {
			// if present in visited
		} else {
			queue = append(queue, Element{value: element.value - 1, count: element.count + 1})
			visited[element.value-1] = element.count + 1
		}
		// Apply Rule : 2
		f := Factors(element.value)
		length := len(f)
		if length > 0 { // if length is not 0 then it is
			var mid int
			if (length % 2) == 0 { // even length
				mid = f[(length)/2]
			} else { // odd length,perfect square
				if length == 1 {
					mid = f[0]
				} else {
					mid = f[(length+1)/2]
				}
			}
			if _, ok := visited[mid]; ok {
				// if it is in the queue
			} else {
				// now mid we will pick and push into the queue
				queue = append(queue, Element{value: mid, count: element.count + 1})
				visited[mid] = element.count + 1
			}
		}
	}
	return 0
}

func main() {

	for i := 2; i < 100; i++ {
		GetMinSteps(i)
	}
}
