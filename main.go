package main

import (
	"flag"
	"fmt"
	"os"

	"ascii/ascii"
)

func main() {
	for _, v := range os.Args {
		if v == "--color" || v == "-color" {
			fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
			return
		}
	}
	var receive ascii.Receiver

	flag.StringVar(&receive.Colorflag, "color", "reset", "color for color input")
	flag.Parse()
	argsPassed := flag.Args()

	msg := receive.SortArg(argsPassed)
	if msg != "" {
		fmt.Println(msg)
		return
	}

	if ascii.IsFlagPassed("color") {
		receive.Color()
	} else {
		receive.ColorCode = ""
		receive.Art()
	}
}
