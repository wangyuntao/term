package term

type Rune = rune

// types: WinResize, Key, Rune
type Event interface {
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
