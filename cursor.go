package term

import (
	"os"
)

var (
	csiCpr = []byte("\x1b[6n")
)

var (
	cursorPosCh = make(chan int)
)

func CursorPosition() (int, int, error) {
	n, err := os.Stdout.Write(csiCpr)
	if err != nil {
		return 0, 0, err
	}

	if n != len(csiCpr) {
		panic("TODO n != len(csiCpr)")
	}

	p := <-cursorPosCh
	return p >> 16 & 0xffff, p & 0xffff, nil
}

func reportCursorPosition(bf []byte) (int, bool) {
	row, col, n, ok := decodeCursorPosition(bf)
	if !ok {
		return 0, false
	}
	cursorPosCh <- (row<<16 | col)
	return n, ok
}

// \x1b[%d;%dR
func decodeCursorPosition(bf []byte) (int, int, int, bool) {
	if len(bf) < 6 || bf[0] != '\x1b' || bf[1] != '[' {
		return 0, 0, 0, false
	}

	row, col := 0, 0
	isRow := true

	// row
	for i := 2; i < len(bf); i++ {
		b := bf[i]
		if isRow {
			if b >= '0' && b <= '9' {
				row = row*10 + int(b-'0')
			} else if b == ';' && row > 0 {
				isRow = false
			} else {
				return 0, 0, 0, false
			}
		} else {
			if b >= '0' && b <= '9' {
				col = col*10 + int(b-'0')
			} else if b == 'R' && col > 0 {
				return row, col, i + 1, true
			} else {
				return 0, 0, 0, false
			}
		}
	}

	return 0, 0, 0, false
}
