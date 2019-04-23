package main

import (
	"fmt"
	"time"

	"github.com/wangyuntao/term"
)

func main() {
	err := term.Init()
	if err != nil {
		panic(err)
	}
	defer term.Cleanup()

	for {
		row, col, err := term.CursorPosition()
		if err != nil {
			panic(err)
		}
		fmt.Println("cursor:", row, col)
		time.Sleep(time.Second * 1)
	}
}
