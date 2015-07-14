package dw

// DoWhile holds the controls for the do...while().
type DoWhile struct {
	end   chan struct{}
	start chan struct{}
	cond  func() bool
}

// Do simulate a do...while().
func Do(hdlr func()) *DoWhile {
	t := &DoWhile{
		end:   make(chan struct{}),
		start: make(chan struct{}),
	}
	go func() {
		<-t.start
		defer close(t.end)
		for {
			hdlr()
			if !t.cond() {
				break
			}
		}
	}()
	return t
}

// While starts the loop with the given condition.
func (dw *DoWhile) While(cond func() bool) {
	dw.cond = cond
	close(dw.start)
	<-dw.end
}
