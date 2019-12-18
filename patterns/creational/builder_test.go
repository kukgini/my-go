package creational

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("NewConcreteBuilder", func() {
	Context("in default context", func() {
		It("should returns non nil", func() {
			builder := NewConcreteBuilder()
			Expect(&builder).NotTo(BeNil())
		})
		It("should set Built to False", func() {
			builder := NewConcreteBuilder()
			Expect(builder.Built).Should(BeFalse())
		})
	})
	Context("after build", func() {
		It("should set Build to True", func() {
			builder := NewConcreteBuilder()
			builder.Build()
			Expect(builder.Built).Should(BeTrue())
		})
	})
	Context("before Build() called", func() {
		It("should returns un-built product", func() {
			builder := NewConcreteBuilder()
			product := builder.GetResult()
			Expect(product.Built).Should(BeFalse())
		})
	})
	Context("after Build() called", func() {
		It("should returns built product", func() {
			builder := NewConcreteBuilder()
			builder.Build()
			product := builder.GetResult()
			Expect(product.Built).Should(BeTrue())
		})
	})
})
