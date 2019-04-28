package promise

import (
	"context"
	"sync"
)

type Promise struct {
	value        interface{}
	error        error
	resolveQueue []ResolveFunc
	rejectQueue  []RejectFunc
	state        State
	cancelFunc   func()
	lock         *sync.Mutex
	waitChannels []chan struct{}
}

func New(executor func(ctx context.Context, resolve ResolveFunc, reject RejectFunc)) Interface {
	promise := &Promise{
		state: StatePending,
		lock:  &sync.Mutex{},
	}
	go promise.run(executor)
	return promise
}

func (p *Promise) run(executor func(ctx context.Context, resolve ResolveFunc, reject RejectFunc)) {
	ctx, cancel := context.WithCancel(context.TODO())
	p.cancelFunc = cancel
	finally := func() {
		p.lock.Lock()
		var waitChannels []chan struct{}
		copy(waitChannels, p.waitChannels)
		p.lock.Unlock()
		for _, ch := range waitChannels {
			ch <- struct{}{}
		}
	}
	resolve := func(result interface{}) {
		p.lock.Lock()
		if p.state != StatePending {
			p.lock.Unlock()
			return
		}
		p.state = StateResolved
		p.value = result
		p.lock.Unlock()
		for _, resolve := range p.resolveQueue {
			resolve(result)
		}
		finally()
	}
	reject := func(err error) {
		p.lock.Lock()
		if p.state != StatePending {
			return
		}
		p.state = StateRejected
		p.error = err
		p.lock.Unlock()
		for _, reject := range p.rejectQueue {
			reject(err)
		}
		finally()
	}

	executor(
		ctx,
		resolve,
		reject)

}

func (p *Promise) resolve(onResolved ThenFunc, onRejected CatchFunc) Interface {
	return New(func(ctx context.Context, onResolveNext ResolveFunc, onRejectNext RejectFunc) {
		resolved := func(value interface{}) {
			if onResolved == nil {
				onResolveNext(value)
			} else {
				nextPromise := onResolved(value)
				if nextPromise != nil {
					nextPromise.
						Then(func(nextValue interface{}) Interface {
							onResolveNext(nextValue)
							return nil
						}).
						Catch(func(err error) Interface {
							onRejectNext(err)
							return nil
						})
				}
			}
		}
		rejected := func(err error) {
			if onRejected == nil {
				onRejectNext(err)
			} else {
				nextPromise := onRejected(err)
				if nextPromise != nil {
					nextPromise.
						Then(func(nextValue interface{}) Interface {
							onResolveNext(nextValue)
							return nil
						}).
						Catch(func(err error) Interface {
							onRejectNext(err)
							return nil
						})
				}
			}
		}
		p.lock.Lock()
		state := p.state
		p.lock.Unlock()
		switch state {
		case StatePending:
			p.resolveQueue = append(p.resolveQueue, resolved)
			p.rejectQueue = append(p.rejectQueue, rejected)
		case StateResolved:
			resolved(p.value)
		case StateRejected:
			rejected(p.error)
		}
	})
}

func (p *Promise) Then(onResolved ThenFunc) Interface {
	return p.resolve(onResolved, nil)
}

func (p *Promise) Catch(onRejected CatchFunc) Interface {
	return p.resolve(nil, onRejected)
}

func (p *Promise) Finally(finallyFunc FinallyFunc) {
	p.Await()
	finallyFunc()
}

func (p *Promise) Abort() {
	p.lock.Lock()
	if p.state != StatePending {
		p.lock.Unlock()
		return
	}
	p.state = StateAborted
	p.lock.Unlock()
	p.cancelFunc()
}

func (p Promise) State() State {
	return p.state
}

func (p Promise) Await() {
	ch := make(chan struct{}, 1)
	p.lock.Lock()
	p.waitChannels = append(p.waitChannels, ch)
	p.lock.Unlock()
	<-ch
}
