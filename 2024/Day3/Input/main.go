package main

import (
	"fmt"
	"os"
)

func readInput() string {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return ""
	}
	return string(data)
}
