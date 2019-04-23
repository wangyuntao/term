package term

import (
	"fmt"
	"os"
	"os/signal"
	"unicode/utf8"

	"golang.org/x/sys/unix"
)

var (
	inputQuit = make(chan int, 2)
)

func initInput(fd int) error {
	err := initKey()
	if err != nil {
		return err
	}

	err = enterRawMode(fd)
	if err != nil {
		return err
	}

	_, err = unix.FcntlInt(uintptr(fd), unix.F_SETFL, unix.O_ASYNC|unix.O_NONBLOCK)
	if err != nil {
		return err
	}

	_, err = unix.FcntlInt(uintptr(fd), unix.F_SETOWN, unix.Getpid())
	if err != nil {
		return err
	}

	go stdinRead(fd)
	return nil
}

func cleanupInput(fd int) {
	inputQuit <- 1
	inputQuit <- 2

	err := exitRawMode(fd)
	if err != nil {
		fmt.Fprintln(os.Stderr, "cleanupInput:", err)
	}
}

func stdinRead(fd int) {
	sigioCh := make(chan os.Signal, 1)
	signal.Notify(sigioCh, unix.SIGIO)

	evtCh := make(chan Event)
	go inputEvt(evtCh)

	bf := make([]byte, 0, 256)
	bfTmp := make([]byte, 64)

	for {
		select {
		case <-sigioCh:
			for {
				n, err := unix.Read(fd, bfTmp)
				if err == unix.EAGAIN || err == unix.EWOULDBLOCK {
					break
				}
				if err != nil {
					panic(err) // TODO handle error properly
				}
				bf = append(bf, bfTmp[:n]...)
			}

			for {
				e, n, ok := decode(bf)
				if n > 0 {
					bf = bf[n:]
				}
				if !ok {
					break
				}
				evtCh <- e
			}
		case <-inputQuit:
			return
		}
	}
}

func decode(bf []byte) (Event, int, bool) {
	if len(bf) == 0 {
		return nil, 0, false
	}

	k, n, ok := decodeKey(bf)
	if ok {
		return k, n, true
	}

	n, ok = reportCursorPosition(bf)
	if ok {
		e, m, ok := decode(bf[n:])
		return e, n + m, ok
	}

	r, n := utf8.DecodeRune(bf)
	if r == utf8.RuneError {
		panic("invalid utf8") // TODO ?
	}
	return r, n, true
}
