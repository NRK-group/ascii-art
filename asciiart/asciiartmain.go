package asciiart

import (
	"asciiart/banner"
	"os"
)

func MainFormatError() string {
	return "Usage: go run . [STRING]\n" + "\n" + "EX: go run . something"
}

func AsciiArtMain() string {
	if len(os.Args[1:]) == 1 {
		return banner.PrintAsciiArt(os.Args[1], banner.Standard)
	} else {
		return MainFormatError()
	}
}
