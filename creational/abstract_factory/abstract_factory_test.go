package abstract_factory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAbstractFactory(t *testing.T) {
	t.Parallel()
	t.Run("", func(t *testing.T) {
		adidasFactory, _ := GetSportsFactory("adidas")
		nikeFactory, _ := GetSportsFactory("nike")

		nikeShoe := nikeFactory.makeShoe()
		nikeShirt := nikeFactory.makeShirt()

		adidasShoe := adidasFactory.makeShoe()
		adidasShirt := adidasFactory.makeShirt()

		assert.Equal(t, adidasShoe.getLogo(), "adidas")
		assert.Equal(t, adidasShirt.getLogo(), "adidas")

		assert.Equal(t, nikeShoe.getLogo(), "nike")
		assert.Equal(t, nikeShirt.getLogo(), "nike")
	})
}
