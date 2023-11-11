package factory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFactory(t *testing.T) {
	t.Parallel()
	t.Run("", func(t *testing.T) {

		ak47, _ := getGun("ak47")
		musket, _ := getGun("musket")

		assert.Equal(t, ak47.getName(), "AK47 gun")
		assert.Equal(t, musket.getName(), "Musket gun")
	})
}
