package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadFile(str string, splitter string) [][]int {
	// fmt.Println("Reading file")
	file, err := os.Open(str)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var res [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		strArr := strings.Split(s, splitter)
		var intArr []int
		for _, val := range strArr {
			intVal, _ := strconv.Atoi(val)
			intArr = append(intArr, intVal)
		}
		res = append(res, intArr)
	}
	return res
}
func main() {
	orderRules := ReadFile("./Input_Day5_Part1.txt", "|")
	order := ReadFile("./Input_Day5_Part2.txt", ",")
	// fmt.Println(orderRules, order)
	sum := 0

	for _, row := range order {
		valid := true

		for {
			swappedValid := true
			for i := 0; i < len(row)-1; i++ {
				if !checkRule(orderRules, row[i], row[i+1]) {
					valid = false
					swappedValid = false
					row = swap(row, i, i+1)
				}
			}
			if swappedValid {
				break
			}
		}

		if !valid {
			sum += row[len(row)/2]
		}
	}
	fmt.Println(sum)
}

func checkRule(arr [][]int, l int, r int) bool {
	for _, row := range arr {
		if row[0] == l && row[1] == r {

			return true
		} else if row[0] == r && row[1] == l {
			return false
		}

	}
	return true
}

func swap(arr []int, i int, j int) []int {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
	return arr
}
