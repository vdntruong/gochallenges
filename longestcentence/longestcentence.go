package longestcentence

import (
	"strings"
)

// func countWords(s string) int {
// 	return len(strings.Fields(s))
// }

// func isInvalid(char rune, list []string) bool {
// 	for _, c := range list {
// 		if c == string(char) {
// 			return true
// 		}
// 	}
// 	return false
// }

// func SolutionT(s string) int {
// 	sentenceList := []string{}
// 	invalidChar := []string{"?", ".", "!"}

// 	curSentence := ""
// 	for _, char := range s {

// 		if isInvalid(char, invalidChar) {
// 			if len(curSentence) != 0 {
// 				sentenceList = append(sentenceList, curSentence)
// 			}
// 			curSentence = ""
// 		}
// 		curSentence += string(char)
// 	}

// 	maxTotalWars := 0
// 	for _, sen := range sentenceList {
// 		logrus.Info("sen ", sen)
// 		if countWords(sen) > maxTotalWars {
// 			maxTotalWars = countWords(sen)
// 		}
// 	}
// 	return maxTotalWars
// }

func Solution(s string) (rs int) {
	sentences := strings.FieldsFunc(s, func(r rune) bool {
		return r == '!' || r == '.' || r == '?'
	})

	for _, sentence := range sentences {
		count := len(strings.Fields(sentence))
		if count > rs {
			rs = count
		}
	}
	return
}

// s = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(s, "?", "."), ".", "!"), "!", "|")
// 	for _, sentence := range strings.Split(s, "|") {
// 		if maxTotalWars < countWords(sentence) {
// 			maxTotalWars = countWords(sentence)
// 		}
// 	}

// maxTotalWars := 0
// 	for _, one := range strings.Split(s, ".") {
// 		for _, two := range strings.Split(one, "?") {
// 			for _, three := range strings.Split(two, "!") {
// 				if maxTotalWars < countWords(three) {
// 					maxTotalWars = countWords(three)
// 				}
// 			}
// 		}
// 	}
// 	return maxTotalWars

// for _, senone := range strings.Split(s, ".") {
// 	if strings.Contains(senone, "?") {
// 		for _, sentwo := range strings.Split(senone, "?") {
// 			if strings.Contains(sentwo, "!") {
// 				for _, senthree := range strings.Split(sentwo, "!") {
// 					if maxTotalWars < countWords(senthree) {
// 						maxTotalWars = countWords(senthree)
// 					}
// 				}
// 			}
// 			if maxTotalWars < countWords(sentwo) {
// 				maxTotalWars = countWords(sentwo)
// 			}
// 		}
// 	}
// 	if maxTotalWars < countWords(senone) {
// 		maxTotalWars += countWords(senone)
// 	}
// }
