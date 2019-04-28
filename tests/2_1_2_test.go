package tests

import (
	"errors"
	"time"

	promise "github.com/joesonw/go-promise"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("2.1.2.1: When resolved, a promise: must not transition to any other state.", func() {
	testResolved(123, func(p promise.Interface, done Done) {
		onResolvedCalled := false

		p.Then(func(interface{}) promise.Interface {
			onResolvedCalled = true
			return nil
		}).Catch(func(error) promise.Interface {
			Expect(onResolvedCalled).To(BeFalse())
			close(done)
			return nil
		})

		time.Sleep(time.Millisecond * 200)
		close(done)
	})

	It("is trying to resolve then immediately reject", func(done Done) {
		d := promise.NewDeferred()
		onResolvedCalled := false

		d.Promise().Then(func(interface{}) promise.Interface {
			onResolvedCalled = true
			return nil
		}).Catch(func(error) promise.Interface {
			Expect(onResolvedCalled).To(BeFalse())
			close(done)
			return nil
		})

		d.Resolve("any")
		d.Reject(errors.New("unknown"))
		time.Sleep(time.Millisecond * 200)
		close(done)
	})

	It("is trying to resolve then reject, delayed", func(done Done) {
		d := promise.NewDeferred()
		onResolvedCalled := false

		d.Promise().Then(func(interface{}) promise.Interface {
			onResolvedCalled = true
			return nil
		}).Catch(func(error) promise.Interface {
			Expect(onResolvedCalled).To(BeFalse())
			close(done)
			return nil
		})

		time.Sleep(time.Millisecond * 50)
		d.Resolve("any")
		d.Reject(errors.New("unknown"))
		time.Sleep(time.Millisecond * 200)
		close(done)
	})

	It("is trying to resolve immediately then reject delayed", func(done Done) {
		d := promise.NewDeferred()
		onResolvedCalled := false

		d.Promise().Then(func(interface{}) promise.Interface {
			onResolvedCalled = true
			return nil
		}).Catch(func(error) promise.Interface {
			Expect(onResolvedCalled).To(BeFalse())
			close(done)
			return nil
		})

		d.Resolve("any")
		time.Sleep(time.Millisecond * 50)
		d.Reject(errors.New("unknown"))
		time.Sleep(time.Millisecond * 200)
		close(done)
	})
})
