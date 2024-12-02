package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Reading file")
	file, err := os.Open("./Input_Day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	validReports := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		strArr := strings.Fields(s)
		var intArr []int
		for _, val := range strArr {
			intVal, _ := strconv.Atoi(val)
			intArr = append(intArr, intVal)
		}
		newArr := removeElement(intArr, 1)
		newArr1 := removeElement(intArr, 0)
		validReport := checkReportValidity(intArr, 0) || checkReportValidity(newArr, 1) || checkReportValidity(newArr1, 1)

		if validReport {
			validReports++
			fmt.Println(s, " is safe")
		} else {
			// fmt.Println(s, " is not valid")
		}
	}
	fmt.Println("Valid Reports=", validReports)
}

func CheckAbsDiff(intArr []int, idx int) bool {
	return math.Abs(float64(intArr[idx-1])-float64(intArr[idx])) >= 1 && math.Abs(float64(intArr[idx-1])-float64(intArr[idx])) < 4
}
func checkReportValidity(intArr []int, dampCount int) bool {
	if dampCount > 1 {
		return false
	}

	if intArr[0] > intArr[1] {
		//Decreasing
		for idx := 1; idx < len(intArr); idx++ {
			if intArr[idx-1] <= intArr[idx] || !(CheckAbsDiff(intArr, idx)) {
				return checkWithDeletions(intArr, dampCount, idx)
			}
		}
	} else {
		//Increasing
		for idx := 1; idx < len(intArr); idx++ {
			if intArr[idx-1] >= intArr[idx] || !(CheckAbsDiff(intArr, idx)) {
				return checkWithDeletions(intArr, dampCount, idx)
			}
		}
	}

	return true
}
func removeElement(originalArray []int, i int) []int {
	newLength := len(originalArray) - 1
	idx := 0
	newArr := make([]int, newLength)
	for index := range originalArray {
		if i != index {
			newArr[idx] = originalArray[index]
			idx++
		}
	}

	return newArr
}

func checkWithDeletions(intArr []int, dampCount int, idx int) bool {
	if dampCount == 0 {
		newArr := removeElement(intArr, idx)
		newArr1 := removeElement(intArr, idx-1)
		return checkReportValidity(newArr, dampCount+1) || checkReportValidity(newArr1, dampCount+1)
	} else {
		return false
	}
}
