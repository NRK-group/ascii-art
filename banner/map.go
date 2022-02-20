package banner

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const Standard string = "banner/standard.txt"

const Shadow string = "banner/shadow.txt"

const Thinkertoy string = "banner/thinkertoy.txt"

func AsciiMap(banner string) map[int][]string {
	file, err := ioutil.ReadFile(banner)
	if err != nil {
		fmt.Println("Invalid file")
	}
	splitFile := strings.Split(string(file), "\n")
	asciiArtMap := make(map[int][]string)
	art := []string{} // slices for every art
	posCount := 0     // position in the map
	ascii := 32       // ascii value counter
	for ascii <= 126 {
		num := (9 * posCount) // equation to get the position in the map
		for i := num; i < num+9; i++ {
			art = append(art, splitFile[i])
		}
		asciiArtMap[ascii] = art
		art = []string{}
		posCount++
		ascii++
	}
	return asciiArtMap
}
