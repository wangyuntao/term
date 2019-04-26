package main

import (
	"github.com/wangyuntao/term"
	"github.com/wangyuntao/terminfo"
)

func main() {
	err := term.Init()
	if err != nil {
		panic(err)
	}
	defer term.Cleanup()

	ti := term.Terminfo()
	ti.Text().Underline().ColorFg(terminfo.ColorCyan).Println("hello")
}
