package creational

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConcreteBuilder_ReturnsNonNil(t *testing.T) {
	t.Parallel()
	builder := NewConcreteBuilder()
	assert.NotNil(t, builder)
}

func TestNewConcreteBuilder_SetsBuildToFalse(t *testing.T) {
	t.Parallel()
	builder := NewConcreteBuilder()
	assert.False(t, builder.built)
}

func TestBuild_SetsBuildToTrue(t *testing.T) {
	t.Parallel()
	builder := NewConcreteBuilder()
	builder.Build()
	assert.True(t, builder.built)
}

func TestGetResult_WhenBuildNotCalled_ReturnsUnBuiltProduct(t *testing.T) {
	t.Parallel()
	builder := NewConcreteBuilder()
	product := builder.GetResult()
	assert.False(t, product.Built)
}

func TestGetResult_WhenBuildCalled_ReturnsBuiltProduct(t *testing.T) {
	t.Parallel()
	builder := NewConcreteBuilder()
	builder.Build()
	product := builder.GetResult()
	assert.True(t, product.Built)
}