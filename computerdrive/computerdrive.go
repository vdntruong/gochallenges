package computerdrive

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	extension = map[string][]string{
		"music": []string{"mp3", "aac", "flac"},
		"image": []string{"jpg", "bmp", "gif"},
		"movie": []string{"mp4", "avi", "mkv"},
	}
)

func getExtension(s string) string {
	arSplit := strings.Split(s, ".")
	return arSplit[len(arSplit)-1]
}
func contain(s []string, someThing string) bool {
	for _, h := range s {
		if h == someThing {
			return true
		}
	}
	return false
}
func kindOfExtension(s string) string {
	for k, ls := range extension {
		if contain(ls, s) {
			return k
		}
	}
	return "other"
}
func getByte(s string) int {
	if i, err := strconv.Atoi(strings.Split(s, "b")[0]); err == nil {
		return i
	}
	return 0
}
func Solution(s string) string {
	arrStr := strings.FieldsFunc(s, func(char rune) bool {
		return char == '\t' || char == ' ' || char == '\n'
	})
	counter := make(map[string]int)
	for i := 0; i < len(arrStr)-1; i += 2 {
		extension := getExtension(arrStr[i])
		kind := kindOfExtension(extension)
		byteN := getByte(arrStr[i+1])
		counter[kind] += byteN
	}
	return fmt.Sprintf("music %db\nimages %db\nmovies %db\nother %db", counter["music"], counter["image"], counter["movie"], counter["other"])
}
