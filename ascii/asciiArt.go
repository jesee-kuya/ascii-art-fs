package ascii

import (
	"fmt"
	"strings"
)

// Ascii prints ASCII art from a given array of characters.
// The characters are extracted from a predefined file array.
// The function takes in four arguments: fileArr (a slice of strings representing the file array),
// wordsArr (a slice of strings representing the words to be printed),
// lettersToColor (a string representing the letters to be colored),
// and color (a string representing the color to be applied).
func Ascii(s Receiver) {
	var count int
	reset := "\033[0m"

	for _, val := range s.WordsArr {
		if val != "" {
			for i := 1; i <= 8; i++ {
				for _, v := range val {
					start := (v - 32) * 9
					if len(s.LettersToColor) == 0 {
						fmt.Print(s.ColorCode + s.FileArr[int(start)+i] + reset)
					} else if strings.Contains(s.LettersToColor, string(v)) {
						fmt.Print(s.ColorCode + s.FileArr[int(start)+i] + reset)
					} else {
						fmt.Print(s.FileArr[int(start)+i])
					}
				}
				fmt.Println()
			}
		} else {
			count++
			if count < len(s.WordsArr) {
				fmt.Println()
			}
		}
	}
}
