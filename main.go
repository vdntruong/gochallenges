package main

import (
	"fmt"
	"sort"
	"strings"
)

func canBeatingWithString(B []string, curRow, curCol, count int, direct string) int {
	fmt.Printf("\nIndex of 0(%v:%v), count %v and direct %s.", curRow, curCol, count, direct)
	if curRow <= 1 {
		return 0
	}

	if curCol > 1 && curCol < len(B)-2 {
		if isX(B, curRow-1, curCol-1) && !isX(B, curRow-2, curCol-2) && isX(B, curRow-1, curCol+1) && !isX(B, curRow-2, curCol+2) {
			return canBeatingWithString(B, curRow-2, curCol-2, count+1, "left") + canBeatingWithString(B, curRow-2, curCol+2, count+1, "right") + 2
		}
	}

	if curCol > 1 {
		if isX(B, curRow-1, curCol-1) && !isX(B, curRow-2, curCol-2) {
			return canBeatingWithString(B, curRow-2, curCol-2, count+1, "left") + 1
		}
	}

	if curCol < len(B)-2 {
		if isX(B, curRow-1, curCol+1) && !isX(B, curRow-2, curCol+2) {
			return canBeatingWithString(B, curRow-2, curCol+2, count+1, "right") + 1
		}
	}

	return 0
}

// SolutionWithString ...
func solutionWithString(B []string) int {
	curRow := 0
	curCol := 0
	for i := 0; i < len(B); i++ {
		if strings.Contains(B[i], `O`) {
			curRow = i
			for j := 0; j < len(B[i]); j++ {
				if string(B[i][j]) == `O` {
					curCol = j
				}
			}
		}
	}
	fmt.Printf("Index of 0(%v:%v) and rs is: ", curRow, curCol)
	return canBeatingWithString(B, curRow, curCol, 0, "strart")
}

func isX(B []string, i, j int) bool {
	return string(B[i][j]) == "X"
}

// Solution ...
func solution(B []string) int {
	curRow := 0
	curCol := 0
	for i := 0; i < len(B); i++ {
		if strings.Contains(B[i], `O`) {
			curRow = i
			for j := 0; j < len(B[i]); j++ {
				if string(B[i][j]) == `O` {
					curCol = j
				}
			}
		}
	}
	fmt.Printf("Index of 0(%v:%v) and rs is: ", curRow, curCol)
	return canBeating(B, curRow, curCol)
}
func canBeating(B []string, curRow, curCol int) int {
	if curRow <= 1 {
		return 0
	}

	if curCol > 1 && curCol < len(B)-2 {
		if isX(B, curRow-1, curCol-1) && !isX(B, curRow-2, curCol-2) && isX(B, curRow-1, curCol+1) && !isX(B, curRow-2, curCol+2) {
			if canBeating(B, curRow-2, curCol-2)+1 > canBeating(B, curRow-2, curCol+2)+1 {
				return canBeating(B, curRow-2, curCol-2) + 1
			}
			return canBeating(B, curRow-2, curCol+2) + 1
		}

	}

	if curCol > 1 {
		if isX(B, curRow-1, curCol-1) && !isX(B, curRow-2, curCol-2) {
			return canBeating(B, curRow-2, curCol-2) + 1
		}
	}

	if curCol < len(B)-2 {
		if isX(B, curRow-1, curCol+1) && !isX(B, curRow-2, curCol+2) {
			return canBeating(B, curRow-2, curCol+2) + 1
		}
	}

	return 0
}

func main() {
	ip := []string{
		`..X...`,
		`......`,
		`....X.`,
		`.X....`,
		`..X.X.`,
		`...O..`,
	}
	ip2 := []string{
		`X....`,
		`.X...`,
		`..O..`,
		`...X.`,
		`.....`,
	}
	ip3 := []string{
		`.....`,
		`.X.X.`,
		`.....`,
		`.X.X.`,
		`..O..`,
	}
	ip4 := []string{
		`.......`,
		`..X.X..`,
		`.......`,
		`X.X.X..`,
		`.......`,
		`..X.X..`,
		`...O...`,
	}
	fmt.Println(solution(ip))
	fmt.Println(solution(ip2))
	fmt.Println(solution(ip3))
	fmt.Println(solution(ip4))
	// fmt.Println(solutionWithString(ip4))
	hello := sort.StringSlice{
		"cba",
		"bca",
		"abc",
	}
	hello.Sort()
	fmt.Println(hello)
}
