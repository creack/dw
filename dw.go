package dw

// DoWhile holds the controls for the do...while().
type DoWhile struct {
	hdlr func()
}

// Do simulate a do...while().
func Do(hdlr func()) *DoWhile {
	return &DoWhile{
		hdlr: hdlr,
	}
}

// While starts the loop with the given condition.
func (dw *DoWhile) While(cond func() bool) {
	for {
		dw.hdlr()
		if !cond() {
			break
		}
	}
}
