package term

import (
	"golang.org/x/sys/unix"
)

var (
	tios *unix.Termios
)

// http://man7.org/linux/man-pages/man3/termios.3.html
func enterRawMode(fd int) error {
	p, err := unix.IoctlGetTermios(fd, unix.TCGETS)
	if err != nil {
		return err
	}

	q := *p   // copy
	tios = &q // save

	p.Iflag &^= unix.IGNBRK | unix.BRKINT | unix.PARMRK | unix.ISTRIP |
		unix.INLCR | unix.IGNCR | unix.ICRNL | unix.IXON

	// p.Oflag &^= unix.OPOST

	p.Lflag &^= unix.ECHO | unix.ECHONL | unix.ICANON |
		//unix.ISIG |
		unix.IEXTEN

	p.Cflag &^= unix.CSIZE | unix.PARENB
	p.Cflag |= unix.CS8

	return unix.IoctlSetTermios(fd, unix.TCSETS, p)
}

func exitRawMode(fd int) error {
	if tios != nil {
		return unix.IoctlSetTermios(fd, unix.TCSETS, tios)
	}
	return nil
}
