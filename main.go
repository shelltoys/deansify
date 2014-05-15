package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"regexp"
)

var version = "0.0.1"
var ansiRegexp = regexp.MustCompile("\x1b[^m]*m")

func stripAnsi(s string) string {
	return ansiRegexp.ReplaceAllLiteralString(s, "")
}

func stripReader(reader *bufio.Reader) {
	for {
		line, err := reader.ReadString('\n')
		if err == nil || err == io.EOF {
			print(stripAnsi(line))
		}

		if err != nil {
			break
		}
	}
}

func stripStdin() {
	reader := bufio.NewReader(os.Stdin)
	stripReader(reader)
}

func stripFile(fileName string) {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatalf("Error opening file %s, %v", fileName, err)
	}

	reader := bufio.NewReader(file)

	stripReader(reader)
}

func main() {
	if len(os.Args) > 1 {
		fileName := os.Args[1]
		stripFile(fileName)
	} else {
		stripStdin()
	}
}
