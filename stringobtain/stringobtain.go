package stringobtain

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

const (
	ADD        = "ADD"
	MOVE       = "MOVE"
	CHANGE     = "CHANGE"
	NOTHING    = "NOTHING"
	IMPOSSIBLE = "IMPOSSIBLE"
)

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
		length := len(T)
		var moveWord string
		canMove := false
		for i := 0; i < length; i++ {
			// consider move case
			if T[i] != S[i] && !canMove {
				moveWord = string(S[i])
				canMove = true
			}

			// change
			if T[i] != S[i] {
				logrus.Info("i ", i)
				logrus.Info("i == length-1", i == length-1)
				if i == length-1 && !canMove {
					return fmt.Sprintf("%s %s %s", CHANGE, string(S[i]), string(T[i]))
				}

				if i < length-1 && canMove && T[i+1] == S[i+1] {
					return fmt.Sprintf("%s %s %s", CHANGE, string(S[i]), string(T[i]))
				}
			}

		}
		return fmt.Sprintf("%s %s", MOVE, moveWord)
	}

	return IMPOSSIBLE
}

// func Solution(S, T string) string {
// 	// nothing
// 	if S == T {
// 		return NOTHING
// 	}
// 	// add, len T = len S +1
// 	if len(S)+1 == len(T) {
// 		for i := 0; i < len(S); i++ {
// 			if T[i] != S[i] {
// 				return fmt.Sprintf("%s %s", ADD, string(T[i]))
// 			}
// 		}
// 		return fmt.Sprintf("%s %s", ADD, string(T[len(T)-1]))
// 	}

// 	// change || move, === len
// 	if len(S) == len(T) {
// 		length := len(T)
// 		flMoved := false
// 		var countChangeWord int
// 		var fistChangeWord string
// 		var secondChangeWord string
// 		var moveWord string

// 		for i := 0; i < length; i++ {
// 			// change
// 			if T[i] != S[i] {
// 				countChangeWord++
// 				if i == length-1 || T[i+1] == S[i+1] {
// 					fistChangeWord = string(S[i])
// 					secondChangeWord = string(T[i])
// 				}
// 				// save fist moved word
// 				if !flMoved {
// 					flMoved = true
// 					moveWord = string(S[i])
// 				}
// 			}
// 			// comfirm move case
// 			if T[length-1-i] != S[length-1-i] {
// 				flMoved = true
// 			}
// 		}
// 		if countChangeWord == 1 {
// 			return fmt.Sprintf("%s %s %s", CHANGE, fistChangeWord, secondChangeWord)
// 		}
// 		return fmt.Sprintf("%s %s", MOVE, moveWord)
// 	}

// 	return IMPOSSIBLE
// }
