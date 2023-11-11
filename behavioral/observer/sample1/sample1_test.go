package sample1

import (
	"testing"
)

func TestObserver1(t *testing.T) {
	t.Parallel()
	t.Run("", func(t *testing.T) {
		stockGrabber := NewStockGrabber()

		obs1 := NewStockObserver("fakeID1", stockGrabber)

		stockGrabber.SetIBMPrice(200)
		stockGrabber.SetAAPLPrice(400)
		stockGrabber.SetGOOGPrice(500)

		obs2 := NewStockObserver("fakeID2", stockGrabber)

		stockGrabber.SetIBMPrice(250)
		stockGrabber.SetAAPLPrice(450)
		stockGrabber.SetGOOGPrice(550)

		stockGrabber.Unregister(obs1)

		stockGrabber.SetIBMPrice(270)
		stockGrabber.SetAAPLPrice(470)
		stockGrabber.SetGOOGPrice(570)

		stockGrabber.Unregister(obs2)
	})
}
