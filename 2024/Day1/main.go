package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Reading file")
	file, err := os.Open("./Input_Day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var leftArr []int
	var rightArr []int
	for scanner.Scan() {
		s := scanner.Text()
		splitStr := strings.Split(s, "   ")
		leftDig, leftErr := strconv.Atoi(splitStr[0])
		if leftErr != nil {
			fmt.Println("Error with Left Dig conversion")
		}
		leftArr = append(leftArr, leftDig)

		rightDig, rightErr := strconv.Atoi(splitStr[1])
		if rightErr != nil {
			fmt.Println("Error with right Dig conversion")
		}
		rightArr = append(rightArr, rightDig)
	}
	sort.Ints(leftArr)
	sort.Ints(rightArr)
	sum := 0
	for i := 0; i < len(leftArr); i++ {
		sum += int(math.Abs(float64(leftArr[i]) - float64(rightArr[i])))
	}
	fmt.Println(sum)

	//Part 2
	//Dictionary of occurences in rightArr
	occurences := make(map[int]int)
	for dig := range rightArr {
		val, ok := occurences[rightArr[dig]]
		// If the key exists
		if ok {
			occurences[rightArr[dig]] = val + 1
		} else {
			occurences[rightArr[dig]] = 1
		}

	}
	similarityScore := 0
	for dig := range leftArr {
		val, ok := occurences[leftArr[dig]]
		if ok {
			similarityScore += leftArr[dig] * val
		}
	}
	fmt.Println("Similarity Score=", similarityScore)

}
