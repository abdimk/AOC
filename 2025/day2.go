package main

import (
	"fmt"
	"strconv"
	"strings"
)


func generateRange(str string)[][]int{
	genRanges := [][]int{}
	ranges := strings.Split(str, ",")
	for _, m := range ranges{
		separe := strings.Split(m, "-")
		max,_ := strconv.Atoi(separe[1])
		min,_ := strconv.Atoi(separe[0])

		genRanges = append(genRanges, []int{min, max})
	}
	return genRanges
}

func checkValid(num int) bool {
    x := strconv.Itoa(num)
    l := len(x)

    for k := 2; k <= l; k++ {
        if l%k == 0 {
            sz := l / k
            sub := x[:sz]
            ok := true

            for i := sz; i < l; i += sz {
                if x[i:i+sz] != sub {
                    ok = false
                    break
                }
            }

            if ok {
                return true
            }
        }
    }
    return false
}


func dayTwo(){
	test := "269194394-269335492,62371645-62509655,958929250-958994165,1336-3155,723925-849457,4416182-4470506,1775759815-1775887457,44422705-44477011,7612653647-7612728309,235784-396818,751-1236,20-36,4-14,9971242-10046246,8796089-8943190,34266-99164,2931385381-2931511480,277-640,894249-1083306,648255-713763,19167863-19202443,62-92,534463-598755,93-196,2276873-2559254,123712-212673,31261442-31408224,421375-503954,8383763979-8383947043,17194-32288,941928989-941964298,3416-9716"
	
	val := generateRange(test)
	total := 0
	for _,v := range val{
		for i:= v[0]; i <= v[1]; i++{
			if checkValid(i){
				total += i;
			}
		}
	}

	fmt.Println(total)


	


	
}


