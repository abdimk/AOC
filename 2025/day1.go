package main

import (
	"fmt"
	"strconv"
	"strings"
	"bufio"
	"errors"
	"os"

)

func readFile(name string)([]string, error){
	file,err := os.Open(name)
	scanned := []string{} 

	defer file.Close()

	if err != nil{
		return []string{}, nil
	}
	
	scanner := bufio.NewScanner(file)

	for scanner.Scan(){
		scanned = append(scanned, scanner.Text())
	}

	return scanned, errors.New("Unable to scan")

}


func main() {
	// rotations := []string{"L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82"}

	rotations, _ := readFile("m.txt")

	password := 0
	password2 := 0
	startingPoint := 50

	
	
	for _,value := range rotations{
		if startingPoint == 0{
			password++
		}

		separteValue := strings.Split(value, "")
		currentPointer := separteValue[0]
		currentNumber, _ := strconv.Atoi(strings.Join(separteValue[1:], ""))
		for range currentNumber{
			switch currentPointer {
				case "L":
					startingPoint = (startingPoint -1 + 100)%100
					
				case "R":
					startingPoint = (startingPoint + 1) % 100
					
			}

			if startingPoint == 0{
				password2++
			}
			
		}
		
	}

	fmt.Println(password)
	fmt.Println(password2)
}