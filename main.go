package main

import (
	"flag"
	"fmt"
	"os"

	"ascii/ascii"
)



// main function is used to read the banner file and print the ascii art
// based on the arguments passed
func main() {
	for _, v := range os.Args {
		if v == "--color" || v == "-color" {
			fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
			return
		}
	}
	var receive ascii.Receiver
	var filename string

	flag.StringVar(&receive.Colorflag, "color", "reset", "color for color input")
	flag.Parse()
	argsPassed := flag.Args()

	if len(argsPassed) == 3 {
		if !ascii.IsFlagPassed("color") {
			fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
			return
		}
		filename = argsPassed[2]
		receive.ArgsPassed = argsPassed[:2]
	} else if len(argsPassed) == 2 {
		if ascii.IsFlagPassed("color") {
			_, err := ascii.GetFileName(argsPassed[1])
			if err != nil {
				filename = "standard"
			} else {
				filename = argsPassed[1]
				receive.ArgsPassed = argsPassed[:1]
			}
		} else {
			filename = argsPassed[1]
			receive.ArgsPassed = argsPassed[:1]
		}
	} else if len(argsPassed) == 1 {
		filename = "standard"
	}

	bannerContent, err := ascii.GetFileName(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	receive.FileArr = bannerContent

	if ascii.IsFlagPassed("color") {
		receive.Color()
	} else {
		receive.ColorCode = ""
		receive.Art()
	}
}
