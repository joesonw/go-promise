package tests

import (
	"errors"

	promise "github.com/joesonw/go-promise"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("2.2.1: Both `onResolved` and `onRejected` are optional arguments.", func() {
	Describe("2.2.1.1: If `onResolved` is not a function, it must be ignored.", func() {
		It("is applied to a directly-rejected promise", func(done Done) {
			c := make(chan error, 0)
			promise.Rejected(errors.New("test error")).Catch(func(err error) promise.Interface {
				c <- err
				return nil
			})
			Expect((<-c).Error()).To(Equal("test error"))
			close(done)
		})
		It("is applied to a promise rejected and then chained off of", func(done Done) {
			c := make(chan error, 0)
			promise.Rejected(errors.New("test error")).Catch(nil).Catch(func(err error) promise.Interface {
				c <- err
				return nil
			})
			Expect((<-c).Error()).To(Equal("test error"))
			close(done)
		})

		It("is applied to a directly-resolved promise", func(done Done) {
			c := make(chan interface{}, 0)
			promise.Resolved(123).Then(func(value interface{}) promise.Interface {
				c <- value
				return nil
			})
			Expect(<-c).To(Equal(123))
			close(done)
		})
		It("is applied to a promise resolved and then chained off of", func(done Done) {
			c := make(chan interface{}, 0)
			promise.Resolved(123).Then(nil).Then(func(value interface{}) promise.Interface {
				c <- value
				return nil
			})
			Expect(<-c).To(Equal(123))
			close(done)
		})
	})
})
