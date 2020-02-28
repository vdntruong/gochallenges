package stringobtain

import (
	"fmt"
	"strings"
)

const (
	ADD        = "ADD"
	MOVE       = "MOVE"
	CHANGE     = "CHANGE"
	NOTHING    = "NOTHING"
	IMPOSSIBLE = "IMPOSSIBLE"
)

func getAllIndex(str string, word string) []int {
	rs := make([]int, 0)
	for i, r := range str {
		if string(r) == word {
			rs = append(rs, i)
		}
	}
	return rs
}

// by at most one simple operation
func Solution(S, T string) string {
	// nothing
	if S == T {
		return NOTHING
	}

	// add, len T = len S +1
	if len(S)+1 == len(T) {
		for i := 0; i < len(S); i++ {
			if T[i] != S[i] {
				return fmt.Sprintf("%s %s", ADD, string(T[i]))
			}
		}
		return fmt.Sprintf("%s %s", ADD, string(T[len(T)-1]))
	}

	// change || move, === len
	if len(S) == len(T) {
		length := len(S)
		for i := 0; i < length; i++ {
			// change
			if T[i] != S[i] && S[:i] == T[:i] {
				if i < length-1 && S[i+1:] == T[i+1:] {
					return fmt.Sprintf("%s %s %s", CHANGE, string(S[i]), string(T[i]))
				}
				if i == length-1 {
					return fmt.Sprintf("%s %s %s", CHANGE, string(S[i]), string(T[i]))
				}
			}

			// move
			if i < length-1 {
				moveWord := string(S[i])
				leaveWords := S[:i] + S[i+1:]

				totalMoveWordNumber := strings.Count(T, moveWord)
				if totalMoveWordNumber == 0 {
					break
				}
				allIndex := getAllIndex(T, moveWord)
				for ; totalMoveWordNumber > 0; totalMoveWordNumber-- {
					j := allIndex[totalMoveWordNumber-1]
					if j == -1 {
						break
					}
					secondLeaveWords := T[:j] + T[j+1:]
					if leaveWords == secondLeaveWords {
						return fmt.Sprintf("%s %s", MOVE, moveWord)
					}
				}

			}
		}
	}
	return IMPOSSIBLE
}
