package promise

import (
	"context"
	"sync"
)

func Race(promises ...Interface) Interface {
	return New(func(ctx context.Context, resolve ResolveFunc, reject RejectFunc) {
		l := sync.Mutex{}
		for i := 0; i < len(promises); i++ {
			promises[i].Then(func(value interface{}) Interface {
				l.Lock()
				defer l.Unlock()
				resolve(value)
				for j := 0; j < len(promises); j++ {
					if j == i {
						continue
					}
					promises[j].Abort()
				}
				return nil
			}).Catch(reject.ToCatch())
		}
	})
}
