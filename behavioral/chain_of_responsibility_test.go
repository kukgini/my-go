package behavioral

import (
	"testing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/kukgini/my-go/behavioral"
)

func TestGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Chain of responsability test suites")
}

var _ = Describe("Behavioral Pattern - Chain of responsability", func() {
	Context("initial state", func() {
		It("should returns something", func() {
			h := NewHandler("Alice", nil, 1)
			Expect(&h).NotTo(BeNil())
		})
		It("should have the name I gave it", func() {
			name := "Alice"
			h := NewHandler(name, nil, 1)
			Expect(h.Name()).To(Equal(name))
		})
		It("should have a next handler", func(){
			next := NewHandler("", nil, 2)
			h := NewHandler("Alice", next, 1)
			Expect(h.Next()).NotTo(BeNil())
		})
		It("should have a handleID", func(){
			handleID := 1234
			h := NewHandler("Alice", nil, handleID)
			Expect(h.HandleID()).To(Equal(handleID))
		})
	})
	Context("when first handler can handle it.", func(){
		It("should be able to handle it", func(){
			a := NewHandler("Alice", nil, 1)
			b := NewHandler("Bob", a, 2)
			expected := "Bob handled 2"
			actual := b.Handle(2)
			Expect(actual).To(Equal(expected))	
		})
	})
	Context("when second handler can handle it.", func(){
		It("should be able to handle it", func(){
			a := NewHandler("Alice", nil, 1)
			b := NewHandler("Bob", a, 2)
			expected := "Alice handled 1"
			actual := b.Handle(1)
			Expect(actual).To(Equal(expected))	
		})
	})
	Context("when any handler can't handle it.", func(){
		It("should be unable to handle it", func(){
			a := NewHandler("Alice", nil, 1)
			b := NewHandler("Bob", a, 2)
			expected := "no handler can handle it."
			actual := b.Handle(0)
			Expect(actual).To(Equal(expected))	
		})
	})
})