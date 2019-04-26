package term

func Do(f func() error) error {
	// setup term
	err := Init()
	if err != nil {
		return err
	}
	defer Cleanup()

	ti := Terminfo()

	// enter ca mode
	err = ti.EnterCaMode()
	if err != nil {
		return err
	}
	defer ti.ExitCaMode()

	// clean screen
	err = ti.ClearScreen()
	if err != nil {
		return err
	}

	return f()
}
