package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	// "sort"
	"strconv"
	"strings"
)

func read(fileName string)(r [][]int, i []int){
	ranges := [][]int{}
	ingredient := []int{}
	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	r1, r2 := 0,0;
	r3 := 0
	for scanner.Scan(){
		r := strings.Split(scanner.Text(), "-")
		if len(r) >= 2{
			// fmt.Println(r)
			r1,_ = strconv.Atoi(r[0])
			r2,_ = strconv.Atoi(r[1])
			ranges = append(ranges, []int{r1, r2})
		}else if r[0] != ""{
			r3,_ = strconv.Atoi(r[0])
			ingredient = append(ingredient, r3)
		}		

	}
	return ranges, ingredient	
}


func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}


func main() {
	p2 := 0

	rag, ing := read("input.txt")

	sort.Slice(rag,func(i, j int) bool {
		return rag[i][0] < rag[j][0]
	})


	// 2, 4, 4 + 6

	current := -1
	for i := range rag {
		// s = rag[i][0]
		// e = rag[i][1]
		// fmt.Println(n[0], n[1], current)
		if current >=  rag[i][0]{
			rag[i][0] = current + 1
			
		}
		if rag[i][0] <= rag[i][1]{
			p2 += rag[i][1] - rag[i][0] + 1
		}
		current = max(current, rag[i][1])

	}
	fmt.Println(p2)

	// fmt.Println(max(12, 100))
	
	countFresh := 0
	for _, r := range ing{
		fresh := false
		for _, m := range rag {
			if r > m[0] && r < m[1]{
				fresh = true
			}
		}
		if fresh{
			countFresh+=1
		}
	}
	fmt.Println(countFresh)

}