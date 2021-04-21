package tools

import "strings"

func Dquote(in string) string {
	return "\"" + in + "\""
}

func FindPos(inString string, searchString string) int {
	pos := strings.Index(inString, searchString)
	//fmt.Println("located", pos, len(inString), searchString, len(searchString))
	return pos
}

func FindStartPos(inString string, searchString string) int {
	pos := FindPos(inString, searchString)
	startPos := pos + len(searchString)
	return startPos
}

//TruncateString truncates a string to a define number of characters
func TruncateString(str string, length int) string {
	if length <= 0 {
		return ""
	}

	// This code cannot support Japanese
	// orgLen := len(str)
	// if orgLen <= length {
	//     return str
	// }
	// return str[:length]

	// Support Japanese
	// Ref: Range loops https://blog.golang.org/strings
	truncated := ""
	count := 0
	for _, char := range str {
		truncated += string(char)
		count++
		if count >= length {
			break
		}
	}
	return truncated
}
