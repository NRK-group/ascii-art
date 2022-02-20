package banner

import (
	"strings"
)

/*
	transform the string to art
*/
func AsciiToArt(arg1, ban string) string {
	art := ""
	p := AsciiMap(ban)
	for i := 1; i <= 8; i++ {
		for _, value := range arg1 {
			art += p[int((value))][i]
		}
		art += "\n"
	}
	return art
}

/*
	conditions
*/

func PrintAsciiArt(arg1, ban string) string {
	art := ""
	switch arg1 {
	case "":
		return ""
	case "\\n":
		return "\n"
	default:
		argSplit := strings.Split(arg1, "\\n")
		for _, word := range argSplit {
			if word != "" {
				art += AsciiToArt(word, ban)
			} else {
				art += "\n"
			}
		}
	}
	return art
}
