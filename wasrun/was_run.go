package wasrun

type WasRun struct {
	RunFlag bool
}

func (w *WasRun) TestMethod() {
	w.RunFlag = true
}
