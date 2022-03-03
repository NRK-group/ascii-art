package banner

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const Standard string = "banner/standard.txt"

// const Shadow string = "banner/shadow.txt"
// const Thinkertoy string = "banner/thinkertoy.txt"

func AsciiMap(banner string) map[rune][]string {
	file, err := ioutil.ReadFile(banner) //read the filex
	if err != nil {
		fmt.Println("Invalid file")
	}
	splitFile := strings.Split(string(file), "\n")
	asciiArtMap := make(map[rune][]string)
	art := []string{} // slices for every art(9 lines)
	posCount := 0     // counter to get position in the text file
	ascii := rune(32) // ascii value counter
	// the loop below loops though the text file and creates a map of slices. The slices
	for ascii <= 126 {
		num := (9 * posCount) // equation to get the position of the art in the text file
		for i := num + 1; i < num+9; i++ {
			art = append(art, splitFile[i]) // this a
		}
		asciiArtMap[ascii] = art
		art = []string{}
		posCount++
		ascii++
	}
	return asciiArtMap
}

/*
	transform the string to art
*/
func AsciiToArt(arg, ban string) string {
	art := ""
	p := AsciiMap(ban)
	for i := 0; i < 8; i++ {
		for _, value := range arg {
			art += p[value][i]
		}
		art += "\n"
	}
	return art
}

/*
	conditions
*/

func PrintAsciiArt(arg, ban string) string {
	art := ""
	switch arg {
	case "":
		return ""
	case "\\n":
		return "\n"
	default:
		argSplit := strings.Split(arg, "\\n")
		for _, word := range argSplit {
			if word == "" {
				art += "\n"
			} else {
				art += AsciiToArt(word, ban)
			}
		}
	}
	return art
}
