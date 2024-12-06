package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInput() [][]rune {
	data, err := os.Open("./Input_Day6.txt")
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

func main() {
	arr := readInput()
	fmt.Println(arr)
	i, j := findInit(arr)
	// fmt.Println(i, j)
	arr[i][j] = 'X'
	dir := "Up"
	for {
		if dir == "Up" {
			if i-1 >= 0 {
				if arr[i-1][j] == '#' {
					dir = turn(dir)
				} else {
					i = i - 1
					arr[i][j] = 'X'
				}
			} else {
				break
			}
		} else if dir == "Down" {
			if i+1 < len(arr) {
				if arr[i+1][j] == '#' {
					dir = turn(dir)
				} else {
					i = i + 1
					arr[i][j] = 'X'
				}
			} else {
				break
			}

		} else if dir == "Left" {
			if j-1 >= 0 {
				if arr[i][j-1] == '#' {
					dir = turn(dir)
				} else {
					j = j - 1
					arr[i][j] = 'X'
				}
			} else {
				break
			}

		} else if dir == "Right" {
			if j+1 < len(arr[0]) {
				if arr[i][j+1] == '#' {
					dir = turn(dir)
				} else {
					j = j + 1
					arr[i][j] = 'X'
				}
			} else {
				break
			}
		}

	}
	//Count 'X'
	fmt.Println(findXCount(arr))

}
func findInit(arr [][]rune) (int, int) {
	for i, row := range arr {
		for j, val := range row {
			if val == '^' {
				return i, j
			}
		}
	}
	return 0, 0
}

func turn(currDir string) string {
	arr := []string{"Up", "Right", "Down", "Left"}
	for idx, val := range arr {
		if val == currDir {
			return arr[(idx+1)%4]
		}
	}
	return ""
}
func findXCount(arr [][]rune) int {
	count := 0
	for _, row := range arr {
		for _, val := range row {
			if val == 'X' {
				count++
			}
		}
	}
	return count
}
