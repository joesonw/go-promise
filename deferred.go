package promise

import (
	"context"
)

type Deferred struct {
	resolve ResolveFunc
	reject  RejectFunc
	promise Interface
	context context.Context
}

func NewDeferred() *Deferred {
	deferred := &Deferred{}
	wait := make(chan struct{}, 1)
	promise := New(func(ctx context.Context, resolve ResolveFunc, reject RejectFunc) {
		wait <- struct{}{}
		deferred.context = ctx
		deferred.resolve = resolve
		deferred.reject = reject
	})
	deferred.promise = promise
	<-wait
	close(wait)
	return deferred
}

func (d *Deferred) Resolve(value interface{}) {
	d.resolve(value)
}

func (d *Deferred) Reject(err error) {
	d.reject(err)
}

func (d *Deferred) Promise() Interface {
	return d.promise
}

func (d *Deferred) Context() context.Context {
	return d.context
}
