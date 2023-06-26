package sample2

import (
	"fmt"
	"github.com/fabriciovz/go_patterns/singleton/sample2/pkg1"
	"github.com/fabriciovz/go_patterns/singleton/sample2/pkg2"
	"github.com/fabriciovz/go_patterns/singleton/sample2/singleton"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInstance(t *testing.T) {
	t.Parallel()
	t.Run("When an instance is called, then it should return a pointer to Singleton", func(t *testing.T) {
		instance := singleton.GetInstance()

		assert.NotNil(t, instance)
	})
	t.Run("When get an instance for the first time, the count should be 0", func(t *testing.T) {
		instance := singleton.GetInstance()

		assert.Equal(t, 0, instance.GetCount())
	})
	t.Run("When AddOne is called for the first time, the count should be 1", func(t *testing.T) {
		instance := singleton.GetInstance()
		instance.AddOne()

		assert.Equal(t, 1, instance.GetCount())
	})
	t.Run("When it gets two instances, the two should have the same reference", func(t *testing.T) {
		instance1 := singleton.GetInstance()
		instance2 := singleton.GetInstance()

		assert.NotNil(t, instance1)
		assert.NotNil(t, instance2)
		assert.Same(t, instance1, instance2)
	})
	t.Run("when AddOne method is called for two different instances, then the two should have the same counter", func(t *testing.T) {
		instance1 := singleton.GetInstance()
		instance2 := singleton.GetInstance()

		instance1.AddOne()
		instance2.AddOne()
		instance2.AddOne()

		assert.Equal(t, instance1.GetCount(), instance2.GetCount())
	})
	t.Run("when they are two singletons from two different packages, then the two should have the same references", func(t *testing.T) {
		instance1 := pkg1.GetInstance1()
		instance2 := pkg2.GetInstance2()

		assert.Same(t, instance1, instance2)
	})
	t.Run("when AddOne method is called from two different pkg instances, then the two should have the same counter", func(t *testing.T) {
		instance1 := pkg1.GetInstance1()
		instance2 := pkg2.GetInstance2()

		instance1.AddOne()
		instance2.AddOne()
		instance2.AddOne()

		fmt.Printf("%p\n", instance1)
		fmt.Printf("%p\n", instance2)

		assert.Equal(t, instance1.GetCount(), instance2.GetCount())
	})
}
