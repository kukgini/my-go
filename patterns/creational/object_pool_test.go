package creational_test

import (
	"testing"

	. "github.com/kukgini/my-go/patterns/creational"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestObjectPool(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Object Pool Test Suites")

}

var _ = Describe("Creational Pattern : Object Pool", func() {
	Context("New Pool", func() {
		It("should not returns nil", func() {
			pool := NewPool(nil)
			Expect(pool).NotTo(BeNil())
		})
	})
})
