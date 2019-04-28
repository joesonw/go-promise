package promise

import "context"

func Resolved(value interface{}) Interface {
	return New(func(_ context.Context, resolve ResolveFunc, _ RejectFunc) {
		resolve(value)
	})
}
