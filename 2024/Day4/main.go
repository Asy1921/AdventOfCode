package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInput() [][]rune {
	data, err := os.Open("./Input_Day4.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return nil
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	arr := make([][]rune, 0)
	i := 0
	for scanner.Scan() {
		s := scanner.Text()
		var row []rune
		for _, val := range s {
			row = append(row, val)
		}
		arr = append(arr, row)
		i++
	}
	return arr
}

var xmas []rune = []rune{'X', 'M', 'A', 'S'}
var countXMAS = 0
var countMAS = 0

func main() {
	arr := readInput()
	for i, row := range arr {
		for j, val := range row {
			if val == 'X' {
				checkXMAS(arr, i, j, 1, "N")
			}
			if val == 'A' {
				checkMAS(arr, i, j)
			}

		}
	}
	fmt.Println(countXMAS)
	fmt.Println(countMAS)
}

func checkXMAS(arr [][]rune, i int, j int, idx int, dir string) {
	if idx == 4 {
		countXMAS++
		// fmt.Println(i, ",", j)
		return
	}
	//Up
	if i-1 >= 0 && arr[i-1][j] == xmas[idx] && (dir == "Up" || dir == "N") {
		checkXMAS(arr, i-1, j, idx+1, "Up")
	}
	//Left
	if j-1 >= 0 && arr[i][j-1] == xmas[idx] && (dir == "Left" || dir == "N") {
		checkXMAS(arr, i, j-1, idx+1, "Left")
	}
	//Down
	if i+1 < len(arr) && arr[i+1][j] == xmas[idx] && (dir == "Down" || dir == "N") {
		checkXMAS(arr, i+1, j, idx+1, "Down")
	}
	//Right
	if j+1 < len(arr[i]) && arr[i][j+1] == xmas[idx] && (dir == "Right" || dir == "N") {
		checkXMAS(arr, i, j+1, idx+1, "Right")
	}
	//UpLeft
	if i-1 >= 0 && j-1 >= 0 && arr[i-1][j-1] == xmas[idx] && (dir == "UpLeft" || dir == "N") {
		checkXMAS(arr, i-1, j-1, idx+1, "UpLeft")
	}
	//UpRight
	if i-1 >= 0 && j+1 < len(arr[i]) && arr[i-1][j+1] == xmas[idx] && (dir == "UpRight" || dir == "N") {
		checkXMAS(arr, i-1, j+1, idx+1, "UpRight")
	}
	//DownLeft
	if i+1 < len(arr) && j-1 >= 0 && arr[i+1][j-1] == xmas[idx] && (dir == "DownLeft" || dir == "N") {
		checkXMAS(arr, i+1, j-1, idx+1, "DownLeft")
	}
	//DownRight
	if i+1 < len(arr) && j+1 < len(arr[i+1]) && arr[i+1][j+1] == xmas[idx] && (dir == "DownRight" || dir == "N") {
		checkXMAS(arr, i+1, j+1, idx+1, "DownRight")
	}

}

func checkMAS(arr [][]rune, i int, j int) {
	l1 := false
	l2 := false

	if i+1 < len(arr) && j+1 < len(arr[i+1]) && i-1 >= 0 && j-1 >= 0 {
		if (arr[i-1][j-1] == 'M' && arr[i+1][j+1] == 'S') || (arr[i-1][j-1] == 'S' && arr[i+1][j+1] == 'M') {
			l1 = true
		}
		if (arr[i-1][j+1] == 'M' && arr[i+1][j-1] == 'S') || (arr[i-1][j+1] == 'S' && arr[i+1][j+-1] == 'M') {
			l2 = true
		}
	}
	if l1 && l2 {
		countMAS++
	}

}
