package term

import "testing"

func TestDecodeCursorPosition(t *testing.T) {
	r, c, n, ok := decodeCursorPosition([]byte("\x1b[2;33R"))
	if !ok {
		t.Fail()
	}

	if r != 2 || c != 33 {
		t.Fail()
	}

	if n != 7 {
		t.Fail()
	}
}
