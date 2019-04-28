package promise

type State int

const (
	StatePending State = iota + 1
	StateResolved
	StateRejected
	StateAborted
)

func (s State) String() string {
	switch s {
	case StatePending:
		return "Pending"
	case StateResolved:
		return "Resolved"
	case StateRejected:
		return "Rejected"
	case StateAborted:
		return "Aborted"
	}
	return ""
}

type RejectFunc func(error)
type ResolveFunc func(interface{})

func (f ResolveFunc) ToThen() ThenFunc {
	return ThenFuncNoReturn(f)
}

func (f RejectFunc) ToCatch() CatchFunc {
	return CatchFuncNoReturn(f)
}

type ThenFunc func(value interface{}) Interface
type CatchFunc func(err error) Interface
type FinallyFunc func()

func ThenFuncNoReturn(f func(interface{})) ThenFunc {
	return func(value interface{}) Interface {
		f(value)
		return nil
	}
}

func CatchFuncNoReturn(f func(err error)) CatchFunc {
	return func(err error) Interface {
		f(err)
		return nil
	}
}

type Interface interface {
	Then(ThenFunc) Interface
	Catch(CatchFunc) Interface
	Finally(FinallyFunc)
	State() State
	Abort()
	Await()
}
