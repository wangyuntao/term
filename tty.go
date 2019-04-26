package term

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"golang.org/x/sys/unix"
)

var (
	ttyIn        int
	ttyOut       *os.File
	ttyOutBuf    *bytes.Buffer
	ttyOutWriter io.Writer
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
	ttyOutWriter = out

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

func Print(a ...interface{}) error {
	_, err := fmt.Fprint(ttyOutWriter, a...)
	return err
}

func Println(a ...interface{}) error {
	_, err := fmt.Fprintln(ttyOutWriter, a...)
	return err
}

func Printf(format string, a ...interface{}) error {
	_, err := fmt.Fprintf(ttyOutWriter, format, a...)
	return err
}

func EnableWriteBuf() {
	if ttyOutBuf == nil {
		ttyOutBuf = bytes.NewBuffer(make([]byte, 0, 256))
		ttyOutWriter = ttyOutBuf
		ti.Writer(ttyOutWriter)
	}
}

func DisableWriteBuf() {
	if ttyOutBuf != nil {
		ttyOutWriter = ttyOut
		ti.Writer(ttyOutWriter)
	}
}

func Flush() error {
	if ttyOutBuf != nil {
		_, err := ttyOut.Write(ttyOutBuf.Bytes())
		if err != nil {
			return err
		}
		ttyOutBuf.Reset()
	}
	return nil
}
