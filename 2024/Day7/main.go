package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var resArr []int = make([]int, 0)
var eqnArr [][]int = make([][]int, 0)

func readInput() {
	data, err := os.Open("./Input_Day7.txt")
	if err != nil {
		fmt.Println("File reading error", err)

	}
	defer data.Close()
	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		s := scanner.Text()
		str := strings.Split(s, ":")
		resVal, _ := strconv.Atoi(str[0])
		resArr = append(resArr, resVal)
		strArr := strings.Fields(str[1])
		var eqnRow []int = make([]int, 0)
		for _, char := range strArr {
			eqnVal, _ := strconv.Atoi(char)
			eqnRow = append(eqnRow, eqnVal)
		}
		eqnArr = append(eqnArr, eqnRow)
	}
	fmt.Println(eqnArr)
}

var validityCheck bool = false

func main() {
	readInput()
	sum := 0
	for i := 0; i < len(resArr); i++ {
		validityCheck = false
		checkEqnValidity(i, 1, eqnArr[i][0])
		if validityCheck {
			sum += resArr[i]
		}
	}
	fmt.Println(sum)
}

func checkEqnValidity(resIdx int, eqnIdx int, curSum int) {
	if eqnIdx == len(eqnArr[resIdx]) {
		if curSum == resArr[resIdx] {
			validityCheck = true
			return
		} else {
			return
		}
	}
	if curSum > resArr[resIdx] {
		return
	}

	checkEqnValidity(resIdx, eqnIdx+1, curSum*eqnArr[resIdx][eqnIdx])
	checkEqnValidity(resIdx, eqnIdx+1, curSum+eqnArr[resIdx][eqnIdx])
	strdig1 := strconv.Itoa(curSum)
	strdig2 := strconv.Itoa(eqnArr[resIdx][eqnIdx])
	sum, _ := strconv.Atoi(strdig1 + strdig2)
	checkEqnValidity(resIdx, eqnIdx+1, sum)
}
