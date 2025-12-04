package main

import (
	"fmt"
	// "reflect"
	// "slices"
	"strconv"
	"strings"
)


func findRepeatedPatternLength(nums []int)int{
	if len(nums) == 0{
		return 0
	}

	key := ""
	for _,n := range nums {
		key += fmt.Sprintf("%d", n)
	}

	for size:= 1; size <= len(key) / 2; size++{
		pattern := key[:size]

		allow := true

		for i := 0; i < len(key); i+= size{
			if key[i:i+size] != pattern {
				allow = false
				break
			}
		}

		if allow && len(key) % size == 0 && len(key)%2 == 0{
			return len(key)
		}
	} 

	return 0
}

func join(num []int)int{
	out := ""

	for _,n := range num{
		out += fmt.Sprintf("%d", n)
	}

	val,_:= strconv.Atoi(out)

	return val
}


func main(){
	m :=  []int{1188511883,1188511884, 1188511885}
	frequency := [][]int{}

	invalidID := []int{}
	
	
	for _,v := range m{
		s := strings.Split(strconv.Itoa(v),"")
		m := []int{}
		for v2 := range len(s){
			c,_ := strconv.Atoi(s[v2])
			m = append(m, c)

		}
		frequency = append(frequency, m)
		// clear(m)
	}

	for _, j := range frequency{
		if out := findRepeatedPatternLength(j); out != 0{
			invalidID = append(invalidID, join(j))
		}
	}


	// find the sum 

	total := 0

	for _,v := range invalidID{
		total+=v
	}

	fmt.Println(total)


	
	


	// output := [][]int{{1,1}, {1,2}, {1, 3}, {1, 4}, {1, 5}, {2, 2}}

}