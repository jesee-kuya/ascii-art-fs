package ascii_test

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"ascii/ascii"
)

// Ascii testcases
var testCases = []struct {
	name                                                   string
	line1, line2, line3, line4, line5, line6, line7, line8 string
	input                                                  string
}{
	{
		"Test1",
		" _              _   _          \n",
		"| |            | | | |         \n",
		"| |__     ___  | | | |   ___   \n",
		"|  _ \\   / _ \\ | | | |  / _ \\  \n",
		"| | | | |  __/ | | | | | (_) | \n",
		"|_| |_|  \\___| |_| |_|  \\___/  \n",
		"                               \n",
		"                               \n",
		"hello",
	},
	{
		"Test2",
		"                                 \n",
		"_|    _|          _| _|          \n",
		"_|    _|   _|_|   _| _|   _|_|   \n",
		"_|_|_|_| _|_|_|_| _| _| _|    _| \n",
		"_|    _| _|       _| _| _|    _| \n",
		"_|    _|   _|_|_| _| _|   _|_|   \n",
		"                                 \n",
		"                                 \n",
		"Hello",
	},
	{
		"Test3",
		"                                                \n",
		"o  o     o o           o-O-o o                o \n",
		"|  |     | |             |   |                | \n",
		"O--O o-o | | o-o         |   O--o o-o o-o o-o o \n",
		"|  | |-' | | | |         |   |  | |-' |   |-'   \n",
		"o  o o-o o o o-o         o   o  o o-o o   o-o O \n",
		"                                                \n",
		"                                                \n",
		"Hello There!",
	},
}

// TestAscii tests the Ascii function
func TestAscii(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// redirect stdout to pipe
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			var fileArr []string

			// read ascii art from a file
			if tc.name == "Test1" {
				fileArr, _ = ascii.Reader("../resources/standard.txt", "\n")
			} else if tc.name == "Test2" {
				fileArr, _ = ascii.Reader("../resources/shadow.txt", "\n")
			} else if tc.name == "Test3" {
				fileArr, _ = ascii.Reader("../resources/thinkertoy.txt", "\r\n")
			}
			// split input into lines
			wordsArr := strings.Split(tc.input, "\n")

			// run the ascii function
			ascii.Ascii(fileArr, wordsArr, "", "")

			// close the write end of pipe
			w.Close()
			// restore stdout
			os.Stdout = old

			// read from the read end of pipe
			var buf bytes.Buffer
			_, err := buf.ReadFrom(r)
			if err != nil {
				t.Fatalf("Error reading from pipe: %v", err)
			}
			// compare output to expected output
			expected := tc.line1 + tc.line2 + tc.line3 + tc.line4 + tc.line5 + tc.line6 + tc.line7 + tc.line8
			if buf.String() != expected {
				t.Errorf("got \n %v expected \n %v", buf.String(), expected)
			}
		})
	}
}
