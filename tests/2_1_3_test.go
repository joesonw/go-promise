package tests

import (
	"errors"
	"time"

	promise "github.com/joesonw/go-promise"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("2.1.3.1: When rejected, a promise: must not transition to any other state.", func() {
	testRejected(errors.New("unknown"), func(p promise.Interface, done Done) {
		onRejectedCalled := false

		p.Then(func(interface{}) promise.Interface {
			Expect(onRejectedCalled).To(BeFalse())
			close(done)
			return nil
		}).Catch(func(error) promise.Interface {
			onRejectedCalled = true
			return nil
		})

		time.Sleep(time.Millisecond * 200)
		close(done)
	})

	It("is trying to reject then immediately resolve", func(done Done) {
		d := promise.NewDeferred()
		onRejectedCalled := false

		d.Promise().Then(func(interface{}) promise.Interface {
			Expect(onRejectedCalled).To(BeFalse())
			close(done)
			return nil
		}).Catch(func(error) promise.Interface {
			onRejectedCalled = true
			return nil

		})

		d.Reject(errors.New("unknown"))
		d.Resolve("any")
		time.Sleep(time.Millisecond * 200)
		close(done)
	})

	It("is trying to reject then resolve, delayed", func(done Done) {
		d := promise.NewDeferred()
		onRejectedCalled := false

		d.Promise().Then(func(interface{}) promise.Interface {
			Expect(onRejectedCalled).To(BeFalse())
			close(done)
			return nil
		}).Catch(func(error) promise.Interface {
			onRejectedCalled = true
			return nil
		})

		time.Sleep(time.Millisecond * 50)
		d.Reject(errors.New("unknown"))
		d.Resolve("any")
		time.Sleep(time.Millisecond * 200)
		close(done)
	})

	It("is trying to reject immediately then resolve delayed", func(done Done) {
		d := promise.NewDeferred()
		onRejectedCalled := false

		d.Promise().Then(func(interface{}) promise.Interface {
			Expect(onRejectedCalled).To(BeFalse())
			close(done)
			return nil
		}).Catch(func(error) promise.Interface {
			onRejectedCalled = true
			return nil
		})

		d.Reject(errors.New("unknown"))
		time.Sleep(time.Millisecond * 50)
		d.Resolve("any")
		time.Sleep(time.Millisecond * 200)
		close(done)
	})
})
