package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("Reading file")
	file, err := os.Open("./Input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	digits := map[string]rune{
		"zero": '0', "one": '1', "two": '2', "three": '3', "four": '4', "five": '5', "six": '6', "seven": '7', "eight": '8', "nine": '9',
	}
	// digitArr := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		// l, r := 0, len(s)-1
		// For Question 1 when only actual digits need to be counted
		// for ; l <= r; l++ {

		// 	if unicode.IsDigit(rune(s[l])) == true {
		// 		break
		// 	}
		// }
		// for ; r >= 0; r-- {

		// 	if unicode.IsDigit(rune(s[r])) == true {
		// 		break
		// 	}
		// }
		// str := string(s[l]) + string(s[r])

		//For question 2 when words also need to be counted
		var digRight rune
		var digLeft rune
		found := false

		for i := 0; i < len(s); i++ {
			splitstr := s[i:]
			if unicode.IsDigit(rune(splitstr[0])) == true {
				digLeft = rune(splitstr[0])
				found = true
				break
			}
			for k := range digits {
				if strings.HasPrefix(splitstr, k) {
					digLeft = rune(digits[k])
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		found = false
		for i := len(s) - 1; i >= 0; i-- {
			splitstr := s[0 : i+1]
			if unicode.IsDigit(rune(splitstr[i])) == true {
				digRight = rune(splitstr[i])
				found = true
				break
			}
			for k := range digits {
				if strings.HasSuffix(splitstr, k) {
					digRight = rune(digits[k])
					found = true
					break

				}
			}
			if found {
				break
			}
		}
		fmt.Println(s)
		fmt.Println(string(digLeft), "+", string(digRight))
		str := string(digLeft) + string(digRight)
		num, err := strconv.Atoi(str)
		if err == nil {
			sum += num
		} else {
			fmt.Println("error")
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(sum)
}
