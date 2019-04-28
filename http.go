package promise

import (
	"context"
	"net/http"
)

func FromHTTP(req *http.Request) Interface {
	return New(func(ctx context.Context, resolve ResolveFunc, reject RejectFunc) {
		res, err := http.DefaultClient.Do(req.WithContext(ctx))
		if err != nil {
			reject(err)
			return
		}
		resolve(res)
	})
}
