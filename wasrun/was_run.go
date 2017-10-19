package wasrun

type WasRun struct {
	RunFlag   bool
	SetUpFlag bool
}

func (w *WasRun) SetUp() {
	w.SetUpFlag = true
}

func (w *WasRun) TestMethod() {
	w.RunFlag = true
}
