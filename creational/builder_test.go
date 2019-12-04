package creational

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"	
	. "github.com/kukgini/my-go/creational"
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
			It("should sets built to False", func(){
				builder := NewConcreteBuilder()
				Expect(builder.Built).Should(BeFalse())
			})
		})
	})
})

// func TestNewConcreteBuilder_ReturnsNonNil(t *testing.T) {
// 	t.Parallel()
// 	builder := NewConcreteBuilder()
// 	assert.NotNil(t, builder)
// }

// func TestNewConcreteBuilder_SetsBuildToFalse(t *testing.T) {
// 	t.Parallel()
// 	builder := NewConcreteBuilder()
// 	assert.False(t, builder.built)
// }

// func TestBuild_SetsBuildToTrue(t *testing.T) {
// 	t.Parallel()
// 	builder := NewConcreteBuilder()
// 	builder.Build()
// 	assert.True(t, builder.built)
// }

// func TestGetResult_WhenBuildNotCalled_ReturnsUnBuiltProduct(t *testing.T) {
// 	t.Parallel()
// 	builder := NewConcreteBuilder()
// 	product := builder.GetResult()
// 	assert.False(t, product.Built)
// }

// func TestGetResult_WhenBuildCalled_ReturnsBuiltProduct(t *testing.T) {
// 	t.Parallel()
// 	builder := NewConcreteBuilder()
// 	builder.Build()
// 	product := builder.GetResult()
// 	assert.True(t, product.Built)
// }