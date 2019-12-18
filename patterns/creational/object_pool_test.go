package creational

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Creational Pattern : Object Pool", func() {
	Context("NewPool", func() {
		It("should not returns nil", func() {
			pool := NewPool(nil)
			Expect(pool).NotTo(BeNil())
		})
		It("should sets new function", func() {
			expectedNew := func() interface{} {
				return float64(10.0)
			}
			pool := NewPool(expectedNew)
			Expect(pool.new).NotTo(BeNil())
		})
		It("should add instance to 'inuse'", func() {
			expectedNew := func() interface{} {
				return float64(10.0)
			}
			pool := NewPool(expectedNew)
			pool.Acquire()
			Expect(len(pool.inuse)).To(Equal(1))
		})
		It("should not add instance to 'available'", func() {
			expectedNew := func() interface{} {
				return float64(10.0)
			}
			pool := NewPool(expectedNew)
			pool.Acquire()
			Expect(len(pool.available)).To(Equal(0))
		})
		It("should use available instance if available", func() {
			expectedNew := func() interface{} {
				return float64(10.0)
			}
			pool := NewPool(expectedNew)
			value := pool.Acquire()
			pool.Release(value)
			pool.Acquire()
			Expect(len(pool.available)).To(Equal(0))
			Expect(len(pool.inuse)).To(Equal(1))
		})
		It("should add instance to available when released", func() {
			expectedNew := func() interface{} {
				return float64(10.0)
			}
			pool := NewPool(expectedNew)
			value := pool.Acquire()
			pool.Release(value)
			Expect(len(pool.available)).To(Equal(1))
			Expect(value).To(Equal(pool.available[0]))
		})
		It("should remove instance from 'inuse'", func() {
			expectedNew := func() interface{} {
				return float64(10.0)
			}
			pool := NewPool(expectedNew)
			value := pool.Acquire()
			pool.Release(value)
			Expect(len(pool.inuse)).To(Equal(0))
		})
	})
})
