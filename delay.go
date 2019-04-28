package promise

import (
	"context"
	"time"
)

func DelayResolved(duration time.Duration, value interface{}) Interface {
	return New(func(_ context.Context, resolve ResolveFunc, _ RejectFunc) {
		time.Sleep(duration)
		resolve(value)
	})
}

func DelayRejected(duration time.Duration, err error) Interface {
	return New(func(_ context.Context, _ ResolveFunc, reject RejectFunc) {
		time.Sleep(duration)
		reject(err)
	})
}
