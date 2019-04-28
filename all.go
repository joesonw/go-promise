package promise

import (
	"context"
	"sync"
)

func All(promises ...Interface) Interface {
	return New(func(ctx context.Context, resolve ResolveFunc, reject RejectFunc) {
		var result []interface{}
		if len(promises) == 0 {
			resolve(result)
			return
		}

		l := sync.Mutex{}
		for i := 0; i < len(promises); i++ {
			promises[i].Then(func(value interface{}) Interface {
				l.Lock()
				defer l.Unlock()
				result = append(result, value)
				if len(result) == len(promises) {
					resolve(result)
				}
				return nil
			}).Catch(reject.ToCatch())
		}
	})
}
