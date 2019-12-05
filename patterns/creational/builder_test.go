package creational_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"	
	. "github.com/kukgini/my-go/patterns/creational"
)

func TestGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Creational Pattern Test Suites")

}

var _ = Describe("Creational Pattern : Builder", func() {
	Describe("NewConcreteBuilder", func() {
		Context("in default context", func() {
			It("should returns non nil", func() {
				builder := NewConcreteBuilder()
				Expect(&builder).NotTo(BeNil())
			})
			It("should set Built to False", func(){
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
})