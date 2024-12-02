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
		validReport := true
		if len(strArr) > 0 {

			if CheckAbsDiff(intArr, 1) {
				if intArr[0] > intArr[1] {
					//Decreasing
					for idx := 1; idx < len(intArr); idx++ {
						if intArr[idx-1] <= intArr[idx] || !(CheckAbsDiff(intArr, idx)) {
							validReport = false
							break
						}
					}
				} else {
					//Increasing
					for idx := 1; idx < len(intArr); idx++ {
						if intArr[idx-1] >= intArr[idx] || !(CheckAbsDiff(intArr, idx)) {
							validReport = false
							break
						}
					}
				}

			} else {

				validReport = false
			}
		}
		if validReport {
			validReports++
			fmt.Println(s, " is valid")
		} else {
			fmt.Println(s, " is not valid")
		}
	}
	fmt.Println("Valid Reports=", validReports)
}

func CheckAbsDiff(intArr []int, idx int) bool {
	return math.Abs(float64(intArr[idx-1])-float64(intArr[idx])) >= 1 && math.Abs(float64(intArr[idx-1])-float64(intArr[idx])) < 4
}
