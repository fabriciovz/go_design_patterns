package state

import (
	"fmt"
	"log"
	"testing"
)

func TestObserver1(t *testing.T) {
	t.Parallel()
	t.Run("", func(t *testing.T) {
		vendingMachine := newVendingMachine(1, 10)

		err := vendingMachine.requestItem()
		if err != nil {
			log.Fatalf(err.Error())
		}

		err = vendingMachine.insertMoney(10)
		if err != nil {
			log.Fatalf(err.Error())
		}

		err = vendingMachine.dispenseItem()
		if err != nil {
			log.Fatalf(err.Error())
		}

		fmt.Println()

		err = vendingMachine.addItem(2)
		if err != nil {
			log.Fatalf(err.Error())
		}

		fmt.Println()

		err = vendingMachine.requestItem()
		if err != nil {
			log.Fatalf(err.Error())
		}

		err = vendingMachine.insertMoney(10)
		if err != nil {
			log.Fatalf(err.Error())
		}

		err = vendingMachine.dispenseItem()
		if err != nil {
			log.Fatalf(err.Error())
		}
	})
}
