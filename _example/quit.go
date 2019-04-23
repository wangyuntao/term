package main

import (
	"github.com/wangyuntao/term"
)

func main() {
	err := term.Init()
	if err != nil {
		panic(err)
	}
	term.Cleanup()
}
