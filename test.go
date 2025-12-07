package main

import (
	"fmt"
	// "slices"
	"sort"
)

func test(){
	m := [][]int{{20, 32}, {54, 81}, {3, 18}, {19, 20}}

	fmt.Println(m)
	
	sort.Slice(m, func(i, j int) bool {
		return m[i][1] < m[j][1]
	})

	fmt.Println(m)
}