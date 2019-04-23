package term

import (
	"os"
	"os/signal"

	"golang.org/x/sys/unix"
)

type WinResize int

var (
	winResizeCh = make(chan os.Signal, 1)
)

func initWin() error {
	signal.Notify(winResizeCh, unix.SIGWINCH)
	return nil
}

func WinSize() (int, int, error) {
	w, err := unix.IoctlGetWinsize(ttyFd, unix.TIOCGWINSZ)
	if err != nil {
		return 0, 0, err
	}
	return int(w.Row), int(w.Col), nil
}
