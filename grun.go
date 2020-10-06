package grun

type TryFunc func(args ...interface{}) []interface{}

type Catchable interface {
	Catch(f func (error))
}

type catcher struct{
	err error
}

func (this catcher) Catch(f func (error)) {
	f(this.err)
}

func Run(f func (TryFunc)) (c Catchable) {
	defer func() {
		if err := recover(); err != nil {
			c = catcher{err.(error)}
		}
	}()
	f(func(args ...interface{}) []interface{} {
		var errObj error = nil
		lastArg := args[len(args)-1]
		switch lastArg.(type){
		case error:
			errObj = lastArg.(error)
		}
		if errObj != nil {
			panic(errObj)
		}
		return args
	})
	return catcher{}
}