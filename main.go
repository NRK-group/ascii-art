package main

import (
	"asciiart/banner"
	"fmt"
	"os"
)

func MainFormatError() string {
	return "Usage: go run . [STRING]\n\nEX: go run . something\n"
}

func AsciiArtMain() string {
	if len(os.Args[1:]) == 1 {
		return banner.PrintAsciiArt(os.Args[1], banner.Standard)
	} else {
		return MainFormatError()
	}
}

func main() {
	fmt.Print(AsciiArtMain())
}
