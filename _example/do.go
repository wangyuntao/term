package main

import (
	"time"

	"github.com/wangyuntao/term"
	"github.com/wangyuntao/terminfo"
)

func main() {
	err := term.Do(func() error {
		ti := term.Terminfo()
		err := ti.Text().Underline().ColorFg(terminfo.ColorMagenta).Println("hello")
		if err != nil {
			return err
		}
		time.Sleep(time.Second * 2)
		return nil
	})

	if err != nil {
		panic(err)
	}
}
