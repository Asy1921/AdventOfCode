package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	fmt.Println("Reading file")
	file, err := os.Open("./Input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		l, r := 0, len(s)-1

		for ; l <= r; l++ {

			if unicode.IsDigit(rune(s[l])) == true {
				break
			}
		}
		for ; r >= 0; r-- {

			if unicode.IsDigit(rune(s[r])) == true {
				break
			}
		}
		str := string(s[l]) + string(s[r])
		num, err := strconv.Atoi(str)
		if err == nil {
			sum += num
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(sum)
}
