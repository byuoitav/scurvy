package scurvy

type RandomDispatcher interface {
	RandomDispatch() (msg string, err error)
}
