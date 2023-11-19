package mediator

import (
	"testing"
)

func TestCommand(t *testing.T) {
	t.Parallel()
	t.Run("", func(t *testing.T) {
		stationManager := newStationManger()

		passengerTrain := &PassengerTrain{
			mediator: stationManager,
		}
		freightTrain := &FreightTrain{
			mediator: stationManager,
		}

		passengerTrain.arrive()
		freightTrain.arrive()
		passengerTrain.depart()
	})
}
