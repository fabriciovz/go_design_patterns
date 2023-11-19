package flyweight

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	t.Parallel()
	t.Run("", func(t *testing.T) {
		game := newGame()

		//Add Terrorist
		game.addTerrorist(TerroristDressType)
		game.addTerrorist(TerroristDressType)
		game.addTerrorist(TerroristDressType)
		game.addTerrorist(TerroristDressType)

		//Add CounterTerrorist
		game.addCounterTerrorist(CounterTerrroristDressType)
		game.addCounterTerrorist(CounterTerrroristDressType)
		game.addCounterTerrorist(CounterTerrroristDressType)

		dressFactoryInstance := getDressFactorySingleInstance()

		for dressType, dress := range dressFactoryInstance.dressMap {
			fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.getColor())
		}
	})
}
