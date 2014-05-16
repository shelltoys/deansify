// Package deansify strips ANSI escape codes from either a named file or
// from STDIN.
package deansify

import (
	"bufio"
	"io"
	"log"
	"os"
	"regexp"
)

var version = "0.0.2"
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

// StripStdin reads text from SDTIN and emits that same text minus any
// ANSI escape codes.
func StripStdin() {
	reader := bufio.NewReader(os.Stdin)
	stripReader(reader)
}

// StripFile reads text from the file located at fileName and emits that same
// text minus any ANSI escape codes.
func StripFile(fileName string) {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatalf("Error opening file %s, %v", fileName, err)
	}

	reader := bufio.NewReader(file)

	stripReader(reader)
}
