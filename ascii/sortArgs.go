package ascii

import "fmt"

func (receive *Receiver) SortArg(argsPassed []string) {
	var filename string
	if len(argsPassed) == 3 {
		if IsFlagPassed("color") {
			fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
			return
		}
		filename = argsPassed[2]
		receive.ArgsPassed = argsPassed[:2]
	} else if len(argsPassed) == 2 {
		if IsFlagPassed("color") {
			_, err := GetFileName(argsPassed[1])
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
		receive.ArgsPassed = argsPassed[:1]
	}

	bannerContent, err := GetFileName(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	receive.FileArr = bannerContent
}
