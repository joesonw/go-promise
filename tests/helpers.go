package tests

import (
	"time"

	promise "github.com/joesonw/go-promise"
	. "github.com/onsi/ginkgo"
)

func testResolved(value interface{}, check func(promise.Interface, Done)) {
	It("should pass already resolved", func(done Done) {
		check(promise.Resolved(value), done)
	})

	It("should pass eventually resolved", func(done Done) {
		check(promise.DelayResolved(time.Millisecond*100, value), done)
	})
}

func testRejected(err error, check func(promise.Interface, Done)) {
	It("should pass already rejected", func(done Done) {
		check(promise.Rejected(err), done)
	})

	It("should pass eventually resolved", func(done Done) {
		check(promise.DelayRejected(time.Millisecond*100, err), done)
	})
}
