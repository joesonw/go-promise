package promise

import (
	"context"
	"sync"
)

func Map(promises map[string]Interface) Interface {
	return New(func(ctx context.Context, resolve ResolveFunc, reject RejectFunc) {
		var result map[string]interface{}
		if len(promises) == 0 {
			resolve(result)
			return
		}

		l := sync.Mutex{}
		for key, p := range promises {
			p.Then(func(value interface{}) Interface {
				l.Lock()
				defer l.Unlock()
				result[key] = value
				if len(result) == len(promises) {
					resolve(result)
				}
				return nil
			}).Catch(reject.ToCatch())
		}
	})
}
