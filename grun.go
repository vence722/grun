package grun

const (
	ErrorNamePanic = "panic"
)

var (
	ErrorNone = CaughtError{}
)

type CaughtError struct {
	Name string
	Err error
	Panic interface{}
}

type ThrowFunc func(name string, err error)

type Catchable interface {
	Catch(f func (CaughtError))
}

type catcher struct{
	caughtError CaughtError
}

func (this catcher) Catch(f func (caughtError CaughtError)) {
	if this.caughtError != ErrorNone {
		f(this.caughtError)
	}
}

func Run(f func (ThrowFunc)) (c Catchable) {
	defer func() {
		if err := recover(); err != nil {
			if caughtErr, ok := err.(CaughtError); ok {
				c = catcher{caughtError: caughtErr}
			} else {
				c = catcher{caughtError: CaughtError{Name: ErrorNamePanic, Panic: err}}
			}
		}
	}()
	f(func(name string, err error) {
		if err != nil {
			panic(CaughtError{Name: name, Err: err})
		}
	})
	return catcher{}
}