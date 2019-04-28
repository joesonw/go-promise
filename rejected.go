package promise

import "context"

func Rejected(err error) Interface {
	return New(func(_ context.Context, _ ResolveFunc, reject RejectFunc) {
		reject(err)
	})
}
