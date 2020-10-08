package grun

type HandleErrFunc func(err error)

type Catchable interface {
	Catch(f func (error))
}

type catcher struct{
	err error
}

func (this catcher) Catch(f func (error)) {
	if this.err != nil {
		f(this.err)
	}
}

func Run(f func (HandleErrFunc)) (c Catchable) {
	defer func() {
		if err := recover(); err != nil {
			c = catcher{err.(error)}
		}
	}()
	f(func(err error) {
		if err != nil {
			panic(err)
		}
	})
	return catcher{}
}