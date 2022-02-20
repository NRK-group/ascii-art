package main

import (
	"io/ioutil"
	"os/exec"
	"strings"
	"testing"
)

func ReadTestFile() []string {
	data, err := ioutil.ReadFile("output.txt")
	if err != nil {
		panic(err)
	}

	contentString := string(data)

	contentSplit := strings.Split(contentString, "#")

	return contentSplit
}

func TestAsciiArt(t *testing.T) {
	testCases := []string{
		"hello",
		"HELLO",
		"HeLlo HuMaN",
		"1Hello 2There",
		"Hello\\nThere",
		"{Hello & There #}",
		"hello There 1 to 2!",
		"MaD3IrA&LiSboN",
		"1a\"#FdwHywR&/()=",
		"{|}~",
		"[\\]^_ 'a",
		"RGB",
		":;<=>?@",
		"\\!\" #$%&'()*+,-./",
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"abcdefghijklmnopqrstuvwxyz",
	}

	expectedOutput := ReadTestFile()

	for index, testCase := range testCases {
		cmd := exec.Command("go", "run", ".", testCase)

		output, _ := cmd.Output()

		if string(output) != expectedOutput[index] {
			t.Errorf("\nTest fails when given case:\n\t\"%s\","+"\nThe test should show:\n%s\nInstead it shows:\n%s", testCase, expectedOutput[index], output)
		}
	}
}
