package bad_singleton_impl

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleton1(t *testing.T) {
	t.Parallel()
	t.Run("When an instance is called, then it should return a pointer to Singleton", func(t *testing.T) {
		var humanService HumanService
		human1 := humanService.GetSingleton()

		assert.NotNil(t, human1)
	})
	t.Run("When get an instance for the first time, the count should be 0", func(t *testing.T) {
		var humanService HumanService

		human1 := humanService.GetSingleton()

		assert.Equal(t, 0, human1.GetAge())
	})
	t.Run("When AddOne is called for the first time, the count should be 1", func(t *testing.T) {
		var humanService HumanService

		human1 := humanService.GetSingleton()
		human1.Birthday()

		assert.Equal(t, 1, human1.GetAge())
	})
	t.Run("When it gets two instances, the two should be equal", func(t *testing.T) {
		var humanService1 HumanService
		var humanService2 HumanService

		human1 := humanService1.GetSingleton()
		human2 := humanService2.GetSingleton()

		assert.NotNil(t, human1)
		assert.NotNil(t, human1)
		assert.Same(t, human1, human2)
	})
	t.Run("when AddOne method is called for two different instances, then the two should have the same counter", func(t *testing.T) {
		var humanService1 HumanService
		var humanService2 HumanService

		human1 := humanService1.GetSingleton()
		human2 := humanService2.GetSingleton()

		human1.Birthday()
		human2.Birthday()
		human2.Birthday()

		fmt.Printf("%p\n", human1)
		fmt.Printf("%p\n", human2)

		assert.Equal(t, human1.GetAge(), human2.GetAge())
	})
}
