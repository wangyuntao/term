package term

import (
	"fmt"
	"os"

	"golang.org/x/sys/unix"
)

var (
	ttyFd int
)

func Init() (err error) {
	defer func() {
		if err != nil {
			Cleanup()
		}
	}()

	ttyFd, err = unix.Open("/dev/tty", unix.O_RDWR, 0)
	if err != nil {
		return err
	}

	err = initWin()
	if err != nil {
		return err
	}

	err = initInput(ttyFd)
	if err != nil {
		return err
	}

	return nil
}

func Cleanup() {
	if ttyFd < 0 {
		return
	}

	cleanupInput(ttyFd)

	err := unix.Close(ttyFd)
	if err != nil {
		fmt.Fprintln(os.Stderr, "cleanup:", err)
	}
}

func PollEvent() Event {
	inputPrepare()

	select {
	case e := <-inputEvtCh:
		return e

	case <-winResizeCh:
		return WinResize(0)
	}
}
