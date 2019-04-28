package tests

//import (
//	"errors"
//	"time"
//
//	"github.com/joesonw/go-promise"
//	. "github.com/onsi/ginkgo"
//	. "github.com/onsi/gomega"
//)
//
//var _ = Describe("2.2.3: If `onRejected` is a function", func() {
//	Describe("2.2.2.1: it must be called after `promise` is rejected, with `promise`’s rejection reason as its first argument", func() {
//		testRejected(errors.New("test"), func(p promise.Interface, done Done) {
//			p.Catch(func(err error) promise.Interface {
//				Expect(err.Error()).To(Equal("test"))
//				close(done)
//				return nil
//			})
//		})
//	})
//
//	Describe("2.2.2.2: it must not be called before `promise` is rejected", func() {
//		It("is rejected after a delay", func() {
//			GinkgoWriter.Write([]byte("2.2.2.2 - I is skipped due to the concurrency nature of go"))
//		})
//		It("never resolved", func() {
//			d := promise.NewDeferred()
//			onRejectCalled := false
//			d.Promise().Catch(func(err error) promise.Interface {
//				onRejectCalled = true
//				return nil
//			})
//
//			time.Sleep(time.Millisecond * 150)
//			Expect(onRejectCalled).To(BeFalse())
//		})
//	})
//
//	Describe("2.2.2.3: it must not be called more than once.", func() {
//		It("is already resolved", func(done Done) {
//			timesCalled := 0
//			promise.Resolved("test").Then(func(value interface{}) promise.Interface {
//				timesCalled++
//				Expect(timesCalled).To(Equal(1))
//				close(done)
//				return nil
//			})
//		})
//
//		It("is trying to resolve a pending promise more than once, immediately", func(done Done) {
//			d := promise.NewDeferred()
//			timesCalled := 0
//			d.Promise().Then(func(value interface{}) promise.Interface {
//				timesCalled++
//				Expect(timesCalled).To(Equal(1))
//				close(done)
//				return nil
//			})
//
//			d.Resolve("test")
//			d.Resolve("123")
//		})
//
//		It("is trying to resovle a pending promise more than once, delayed", func(done Done) {
//			d := promise.NewDeferred()
//			timesCalled := 0
//			d.Promise().Then(func(value interface{}) promise.Interface {
//				timesCalled++
//				Expect(timesCalled).To(Equal(1))
//				close(done)
//				return nil
//			})
//
//			time.Sleep(time.Millisecond * 50)
//			d.Resolve("test")
//			d.Resolve("123")
//		})
//
//		It("is trying to resovle a pending promise more than once, immediately then delayed", func(done Done) {
//			d := promise.NewDeferred()
//			timesCalled := 0
//			d.Promise().Then(func(value interface{}) promise.Interface {
//				timesCalled++
//				Expect(timesCalled).To(Equal(1))
//				close(done)
//				return nil
//			})
//
//			d.Resolve("test")
//			time.Sleep(time.Millisecond * 50)
//			d.Resolve("123")
//		})
//
//		It("when multiple `then` calls are made, spaced aprt in time", func(done Done) {
//			d := promise.NewDeferred()
//			timesCalled := []int{0, 0, 0}
//
//			d.Promise().Then(func(value interface{}) promise.Interface {
//				timesCalled[0]++
//				Expect(timesCalled[0]).To(Equal(1))
//				return nil
//			})
//
//			go func() {
//				time.Sleep(time.Millisecond * 50)
//				d.Promise().Then(func(value interface{}) promise.Interface {
//					timesCalled[1]++
//					Expect(timesCalled[1]).To(Equal(1))
//					return nil
//				})
//			}()
//
//			go func() {
//				time.Sleep(time.Millisecond * 100)
//				d.Promise().Then(func(value interface{}) promise.Interface {
//					timesCalled[2]++
//					Expect(timesCalled[2]).To(Equal(1))
//					close(done)
//					return nil
//				})
//			}()
//
//			time.Sleep(time.Millisecond * 150)
//			d.Resolve("test")
//		})
//
//		It("whe `then` is interleaved with resolve", func(done Done) {
//			d := promise.NewDeferred()
//			timesCalled := []int{0, 0}
//
//			d.Promise().Then(func(value interface{}) promise.Interface {
//				timesCalled[0]++
//				Expect(timesCalled[0]).To(Equal(1))
//				return nil
//			})
//			d.Resolve("test")
//			d.Promise().Then(func(value interface{}) promise.Interface {
//				timesCalled[1]++
//				Expect(timesCalled[1]).To(Equal(1))
//				close(done)
//				return nil
//			})
//		})
//	})
//})
