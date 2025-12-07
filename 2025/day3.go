package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	// "sort"
	"strconv"
	"strings"
)



func largest(num string) int{
	array := []int{}
	sp := strings.Split(num, "")
	for _,n:= range sp{
		ok, _ := strconv.Atoi(n)
		array = append(array, ok)
	}
	// n := ""

	if len(array) < 2 {
		return 0
	}

	maxResult := 0
	
	for i := 0; i < len(array) -1; i++{
		maxr := 0

		for j := i+1; j < len(array); j++{
			if array[j] > maxr {
				maxr = array[j]
			}
		}
		pair := array[i]*10 + maxr

		if pair > maxResult {
			maxResult = pair
		}
	}

	return maxResult



	
	
}
func partOne() {
	file, err := os.Open("m.txt")

	if err != nil {
		log.Fatalf("Error opening file :%v", err)
	}
	defer file.Close()

	total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		total += largest(scanner.Text())
	}

	fmt.Println(total)

}