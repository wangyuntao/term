package term

func Init() (err error) {
	defer func() {
		if err != nil {
			Cleanup()
		}
	}()

	err = initTty()
	if err != nil {
		return err
	}

	err = initWin()
	if err != nil {
		return err
	}

	err = initInput(ttyIn)
	if err != nil {
		return err
	}

	return nil
}

func Cleanup() {
	cleanupInput(ttyIn)
	cleanupWin()
	cleanupTty()
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
