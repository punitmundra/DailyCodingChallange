package main

import (
	"fmt"
	"sort"
)

type int2DSlice [][]int

func (c int2DSlice) Len() int      { return len(c) }
func (c int2DSlice) Swap(i, j int) { c[i], c[j] = c[j], c[i] }

// Compare 2nd index
func (c int2DSlice) Less(i, j int) bool { return c[i][1] < c[j][1] }

//=========================================================================//

var sortColumnIndex int

// sortByIndexAscending sorts two-dimensional int in an ascending order, at a specified index.
type sortByIndexAscending [][]int

func (s sortByIndexAscending) Len() int {
	return len(s)
}

func (s sortByIndexAscending) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortByIndexAscending) Less(i, j int) bool {
	return s[i][sortColumnIndex] < s[j][sortColumnIndex]
}

// intAscending sorts two dimensional int in a descending order.
func intAscending(rows [][]int, idx int) [][]int {
	sortColumnIndex = idx
	sort.Sort(sortByIndexAscending(rows))
	return rows
}

//=========================================================================//
type sortByIndexDescending [][]int

func (s sortByIndexDescending) Len() int {
	return len(s)
}

func (s sortByIndexDescending) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortByIndexDescending) Less(i, j int) bool {
	return s[i][sortColumnIndex] > s[j][sortColumnIndex]
}

// intDescending sorts two dimensional int in a descending order.
func intDescending(rows [][]int, idx int) [][]int {
	sortColumnIndex = idx
	sort.Sort(sortByIndexDescending(rows))
	return rows
}

func main() {

	c := int2DSlice{{2, 3},
		{7, 2},
		{2, 8},
		{5, 4},
	}
	sort.Sort(c)
	fmt.Printf("%v\n", c)

	rows := [][]int{
		[]int{2, 3},
		[]int{4, 7},
		[]int{9, 7},
		[]int{6, 6},
		[]int{1, 2},
		[]int{2, 3},
		[]int{4, 3},
		[]int{4, 5},
		[]int{10, 8},
		[]int{8, 1},
	}

	rs1 := intAscending(rows, 0)
	fmt.Println("rs1:", rs1)

	rs2 := intAscending(rows, 1)
	fmt.Println("rs2:", rs2)

	rs3 := intDescending(rows, 0)
	fmt.Println("rs3:", rs3)

	rs4 := intDescending(rows, 1)
	fmt.Println("rs4:", rs4)

}
