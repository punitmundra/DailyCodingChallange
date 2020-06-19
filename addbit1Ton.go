/*

Write an algorithm that finds the total number of set bits in all integers between 1 and N.
*/

package main

import "fmt"

func CountBit(start, end int, ch chan int) {
	var count int
	sum := 0
	for i := start; i <= end; i++ {
		x := i
		for count = 0; x != 0; count++ {
			x &= x - 1
		}
		sum += count
	}
	ch <- sum
}

var m map[interface{}]intference{}

func main() {
	fmt.Println(" Enter the Number:")
	ch := make(chan int, 2)
	var n int
	fmt.Scanf("%d", &n)
	go CountBit(1, n/2, ch)
	go CountBit(n/2+1, n, ch)
	c := <-ch + <-ch
  fmt.Println(c)
	}
