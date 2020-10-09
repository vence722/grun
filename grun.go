package grun

var (
	ErrorNone = CaughtError{}
)

type CaughtError struct {
	Name string
	Err error
}

type HandleErrFunc func(name string, err error)

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

func Run(f func (HandleErrFunc)) (c Catchable) {
	defer func() {
		if err := recover(); err != nil {
			c = catcher{caughtError: err.(CaughtError)}
		}
	}()
	f(func(name string, err error) {
		if err != nil {
			panic(CaughtError{Name: name, Err: err})
		}
	})
	return catcher{}
}