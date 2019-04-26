package term

import (
	"io"

	"github.com/wangyuntao/terminfo"
)

var (
	ti *terminfo.Terminfo
)

func initTerminfo(w io.Writer) error {
	i, err := terminfo.LoadEnv()
	if err != nil {
		return err
	}
	i.Writer(w)
	ti = i
	return nil
}

func cleanupTerminfo() {
	ti.WriterRestore()
}

func Terminfo() *terminfo.Terminfo {
	return ti
}
