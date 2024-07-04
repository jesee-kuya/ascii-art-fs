package ascii

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"os"
	"strings"
)

// Reader reads the banner file and splits it into a slice of strings
// using the separator provided. It returns an error and slice of strings
func Reader(filename string, sepp string) ([]string, error) {
	standardSum := "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf"
	shadowSum := "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73"
	thinkerSum := "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3"

	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, errors.New("error reading banner file")
	}

	// check if the bannerfile has all the values
	h := sha256.New()
	h.Write(file)
	sum := h.Sum(nil)
	hash := fmt.Sprintf("%x", sum)

	if hash != standardSum && hash != shadowSum && hash != thinkerSum {
		return nil, fmt.Errorf(" Error: The content of the banner file has been modified")
	}

	content := strings.Split(string(file), sepp)
	return content, nil
}
