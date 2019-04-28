package main

import (
	"fmt"

	"github.com/wangyuntao/term"
)

func main() {
	err := term.Init()
	if err != nil {
		panic(err)
	}
	defer term.Cleanup()

	for {
		e := term.PollEvent()

		switch v := e.(type) {
		case term.Key:
			fmt.Println("Key:", v)

		case term.Rune:
			fmt.Printf("Rune: %d\n", v)

		case term.AltKey:
			fmt.Println("AltKey:", v)

		case term.AltRune:
			fmt.Printf("AltRune: %d\n", v)

		case term.WinResize:
			row, col, err := term.WinSize()
			if err != nil {
				panic(err)
			}
			fmt.Println("winResize:", row, col)

		}
	}
}
