package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readInput() string {
	data, err := os.ReadFile("./Input_Day3.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return ""
	}
	return string(data)
}
func manageSKips(input string) string {
	newStr := ""
	copy := true
	for idx := range input {
		if strings.HasPrefix(input[idx:], "don't()") {
			copy = false
		}
		if strings.HasPrefix(input[idx:], "do()") {
			copy = true
		}
		if copy {
			newStr += string(input[idx])
		}
	}
	return newStr
}

func main() {
	inputStr := readInput()
	sum := calcSum(inputStr)
	fmt.Println("Part1:", sum)
	newStr := manageSKips(inputStr)
	sum1 := calcSum(newStr)
	fmt.Println("Part2", sum1)

}
func calcSum(inputStr string) int {
	digits := regexp.MustCompile(`mul\((\d+),(\d+)\)`).FindAllStringSubmatchIndex(inputStr, -1)
	sum := 0
	for _, arr := range digits {
		dig1, err := strconv.Atoi(inputStr[arr[2]:arr[3]])
		if err != nil {
			continue
		}
		dig2, err := strconv.Atoi(inputStr[arr[4]:arr[5]])
		if err != nil {
			continue
		}
		sum += dig1 * dig2

	}
	return sum
}
