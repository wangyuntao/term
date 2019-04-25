package term

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/sys/unix"
)

var (
	ttyIn  int
	ttyOut *os.File
)

func initTty() error {
	in, err := unix.Open("/dev/tty", unix.O_RDONLY, 0)
	if err != nil {
		return err
	}
	ttyIn = in

	out, err := os.OpenFile("/dev/tty", os.O_WRONLY, 0)
	if err != nil {
		return err
	}
	ttyOut = out

	return nil
}

func cleanupTty() {
	if ttyIn > 0 {
		err := unix.Close(ttyIn)
		if err != nil {
			fmt.Fprintln(os.Stderr, "cleanupTty:", err)
		}
	}

	if ttyOut != nil {
		err := ttyOut.Close()
		if err != nil {
			fmt.Fprintln(os.Stderr, "cleanupTty:", err)
		}
	}
}

func Writer() io.Writer {
	return ttyOut
}

func Print(a ...interface{}) (int, error) {
	return fmt.Fprint(ttyOut, a...)
}

func Println(a ...interface{}) (int, error) {
	return fmt.Fprintln(ttyOut, a...)
}

func Printf(format string, a ...interface{}) (int, error) {
	return fmt.Fprintf(ttyOut, format, a...)
}
