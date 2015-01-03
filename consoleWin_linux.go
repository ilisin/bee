package main

import (
	"fmt"
)

func ConsoleWinOut(level int, text string) {
	fmt.Println(text)
}

func ConsoleOutWithLinuxFmt(text string) {
	fmt.Print(text)
}
